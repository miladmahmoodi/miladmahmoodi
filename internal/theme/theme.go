package theme

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"go.yaml.in/yaml/v4"

	"github.com/miladmahmoodi/miladmahmoodi/internal/config"
)

// Theme is the interface every Forge theme must satisfy.
type Theme interface {
	Name() string
	Render(data RenderData) (string, error)
}

// Meta holds the metadata declared in a theme's theme.yml.
type Meta struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Author      string `yaml:"author"`
	Version     string `yaml:"version"`
}

// RenderData is the full context passed into every template.
type RenderData struct {
	Config      *config.Config
	GeneratedAt string
	Version     string
	Sections    map[string]string
}

// fsTheme is the concrete implementation backed by an fs.FS.
type fsTheme struct {
	meta      Meta
	templates *template.Template
}

func (t *fsTheme) Name() string { return t.meta.Name }

func (t *fsTheme) Render(data RenderData) (string, error) {
	var buf strings.Builder
	if err := t.templates.ExecuteTemplate(&buf, "base.md.tmpl", data); err != nil {
		return "", fmt.Errorf("rendering theme %q: %w", t.meta.Name, err)
	}
	return buf.String(), nil
}

// Load resolves a theme by name. It first looks in the local ./themes/<name>/
// directory; if not found it falls back to the embedded FS.
func Load(name string, embedded fs.FS) (Theme, error) {
	// Prefer local override (useful for custom themes and theme authoring).
	local := filepath.Join("themes", name)
	if _, err := os.Stat(local); err == nil {
		return loadFromOS(name, os.DirFS(local))
	}

	sub, err := fs.Sub(embedded, filepath.Join("themes", name))
	if err != nil {
		return nil, fmt.Errorf("theme %q not found in embedded assets: %w", name, err)
	}
	return loadFromOS(name, sub)
}

func loadFromOS(name string, themeFS fs.FS) (Theme, error) {
	metaBytes, err := fs.ReadFile(themeFS, "theme.yml")
	if err != nil {
		return nil, fmt.Errorf("theme %q: missing theme.yml: %w", name, err)
	}

	var meta Meta
	if err := yaml.Unmarshal(metaBytes, &meta); err != nil {
		return nil, fmt.Errorf("theme %q: parsing theme.yml: %w", name, err)
	}

	funcMap := template.FuncMap{
		"now": func() string { return time.Now().Format("Mon Jan 02 2006") },
		"add": func(a, b int) int { return a + b },
		"join": func(sep string, items []string) string {
			result := ""
			for i, item := range items {
				if i > 0 {
					result += sep
				}
				result += item
			}
			return result
		},
		"pad": func(width int, s string) string {
			for len(s) < width {
				s += " "
			}
			return s
		},
		"trim": strings.TrimSpace,
		// maxCatLen returns the length of the longest category name + 2 for spacing.
		"maxCatLen": func(groups []config.SkillGroup) int {
			max := 0
			for _, g := range groups {
				if len(g.Category) > max {
					max = len(g.Category)
				}
			}
			return max + 2
		},
	}

	tmpl, err := template.New("").Funcs(funcMap).ParseFS(themeFS, "templates/*.md.tmpl")
	if err != nil {
		return nil, fmt.Errorf("theme %q: parsing templates: %w", name, err)
	}

	return &fsTheme{meta: meta, templates: tmpl}, nil
}

// NewRenderData builds the context for template execution.
func NewRenderData(cfg *config.Config, version string) RenderData {
	return RenderData{
		Config:      cfg,
		GeneratedAt: time.Now().Format("Mon Jan 02 2006 15:04:05 MST"),
		Version:     version,
	}
}
