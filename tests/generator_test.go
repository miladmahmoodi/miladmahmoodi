package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/miladmahmoodi/forge/internal/generator"
)

// minimalThemeFS builds an in-memory FS with a minimal terminal theme for tests.
func minimalThemeFS() fstest.MapFS {
	return fstest.MapFS{
		"themes/terminal/theme.yml": &fstest.MapFile{
			Data: []byte("name: terminal\ndescription: test\nauthor: test\nversion: 0.0.0\n"),
		},
		"themes/terminal/templates/base.md.tmpl": &fstest.MapFile{
			Data: []byte("# {{.Config.Name}}\n{{.Config.Role}}\nv{{.Version}}\n"),
		},
	}
}

func TestGenerate_BasicOutput(t *testing.T) {
	yml := `
name:     "Alice"
username: "alice"
role:     "Engineer"
theme:    "terminal"
`
	configPath := writeTempConfig(t, yml)
	outPath := filepath.Join(t.TempDir(), "README.md")

	result, err := generator.Generate(generator.Options{
		ConfigPath: configPath,
		OutputPath: outPath,
		ThemesFS:   minimalThemeFS(),
	})
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}

	if result.ThemeName != "terminal" {
		t.Errorf("ThemeName = %q, want %q", result.ThemeName, "terminal")
	}
	if result.BytesOut == 0 {
		t.Error("expected non-empty output")
	}

	content, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("reading output: %v", err)
	}
	if !strings.Contains(string(content), "Alice") {
		t.Error("output does not contain expected name")
	}
	if !strings.Contains(string(content), "Engineer") {
		t.Error("output does not contain expected role")
	}
}

func TestGenerate_DryRun(t *testing.T) {
	yml := `
name:     "Bob"
username: "bob"
role:     "Engineer"
`
	configPath := writeTempConfig(t, yml)

	result, err := generator.Generate(generator.Options{
		ConfigPath: configPath,
		ThemesFS:   minimalThemeFS(),
		DryRun:     true,
	})
	if err != nil {
		t.Fatalf("Generate() dry-run error: %v", err)
	}
	if result.OutputPath != "(dry-run)" {
		t.Errorf("expected dry-run output path, got %q", result.OutputPath)
	}
}

func TestGenerate_InvalidConfig(t *testing.T) {
	yml := `name: ""`
	configPath := writeTempConfig(t, yml)

	_, err := generator.Generate(generator.Options{
		ConfigPath: configPath,
		ThemesFS:   minimalThemeFS(),
	})
	if err == nil {
		t.Error("expected error for invalid config, got nil")
	}
}

func TestGenerate_UnknownTheme(t *testing.T) {
	yml := `
name:     "Carol"
username: "carol"
role:     "Engineer"
theme:    "nonexistent-theme"
`
	configPath := writeTempConfig(t, yml)

	_, err := generator.Generate(generator.Options{
		ConfigPath: configPath,
		ThemesFS:   minimalThemeFS(),
	})
	if err == nil {
		t.Error("expected error for unknown theme, got nil")
	}
}

func TestGenerate_OutputFileIsWritten(t *testing.T) {
	yml := `
name:     "Dave"
username: "dave"
role:     "Engineer"
`
	configPath := writeTempConfig(t, yml)
	outPath := filepath.Join(t.TempDir(), "out.md")

	_, err := generator.Generate(generator.Options{
		ConfigPath: configPath,
		OutputPath: outPath,
		ThemesFS:   minimalThemeFS(),
	})
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}

	info, err := os.Stat(outPath)
	if err != nil {
		t.Fatalf("output file not created: %v", err)
	}
	if info.Size() == 0 {
		t.Error("output file is empty")
	}
}
