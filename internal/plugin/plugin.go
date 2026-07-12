package plugin

import (
	"fmt"

	"github.com/miladmahmoodi/forge/internal/config"
)

// Plugin is the interface every Forge plugin must implement.
// Plugins produce a markdown string that gets injected into the rendered profile.
type Plugin interface {
	// Name returns the unique plugin identifier used in config.yml.
	Name() string

	// Render produces the markdown output for this plugin.
	// The returned string is injected at the plugin's designated slot in the template.
	Render(cfg *config.Config, opts map[string]any) (string, error)
}

// Result holds the output of a single plugin execution.
type Result struct {
	PluginName string
	Output     string
	Err        error
}

// RunAll executes all enabled plugins and returns their results.
func RunAll(cfg *config.Config, registry map[string]Plugin) []Result {
	var results []Result

	for _, entry := range cfg.Plugins {
		if !entry.Enabled {
			continue
		}

		p, ok := registry[entry.Name]
		if !ok {
			results = append(results, Result{
				PluginName: entry.Name,
				Err:        fmt.Errorf("unknown plugin %q", entry.Name),
			})
			continue
		}

		output, err := p.Render(cfg, entry.Options)
		results = append(results, Result{
			PluginName: entry.Name,
			Output:     output,
			Err:        err,
		})
	}

	return results
}
