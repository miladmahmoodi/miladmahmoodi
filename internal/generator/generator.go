package generator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/miladmahmoodi/forge/internal/config"
	"github.com/miladmahmoodi/forge/internal/plugin"
	"github.com/miladmahmoodi/forge/internal/theme"
)

const forgeVersion = "0.1.0"

// Options controls the behaviour of a single generation run.
type Options struct {
	ConfigPath string
	OutputPath string
	ThemesFS   fs.FS
	Plugins    map[string]plugin.Plugin
	DryRun     bool
}

// Result holds the outcome of a generation run.
type Result struct {
	ConfigPath string
	OutputPath string
	ThemeName  string
	BytesOut   int
}

// Generate executes the full pipeline:
//
//	config.yml → Load → Validate → Theme → Plugins → Render → README.md
func Generate(opts Options) (*Result, error) {
	cfg, err := config.Load(opts.ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	if errs := config.Validate(cfg); errs.HasErrors() {
		return nil, fmt.Errorf("invalid config:\n%s", errs)
	}

	t, err := theme.Load(cfg.Theme, opts.ThemesFS)
	if err != nil {
		return nil, fmt.Errorf("load theme: %w", err)
	}

	pluginRegistry := opts.Plugins
	if pluginRegistry == nil {
		pluginRegistry = plugin.DefaultRegistry()
	}
	pluginResults := plugin.RunAll(cfg, pluginRegistry)
	for _, pr := range pluginResults {
		if pr.Err != nil {
			fmt.Fprintf(os.Stderr, "  [warn] plugin %q: %v\n", pr.PluginName, pr.Err)
		}
	}

	data := theme.NewRenderData(cfg, forgeVersion)

	output, err := t.Render(data)
	if err != nil {
		return nil, fmt.Errorf("render: %w", err)
	}

	if opts.DryRun {
		fmt.Print(output)
		return &Result{
			ConfigPath: opts.ConfigPath,
			OutputPath: "(dry-run)",
			ThemeName:  cfg.Theme,
			BytesOut:   len(output),
		}, nil
	}

	outPath := opts.OutputPath
	if outPath == "" {
		outPath = filepath.Join(filepath.Dir(opts.ConfigPath), "README.md")
	}

	if err := os.WriteFile(outPath, []byte(output), 0o644); err != nil {
		return nil, fmt.Errorf("writing %s: %w", outPath, err)
	}

	return &Result{
		ConfigPath: opts.ConfigPath,
		OutputPath: outPath,
		ThemeName:  cfg.Theme,
		BytesOut:   len(output),
	}, nil
}
