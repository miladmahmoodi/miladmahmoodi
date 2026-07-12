package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/miladmahmoodi/forge/internal/config"
)

func TestLoad_ValidConfig(t *testing.T) {
	yml := `
name:     "Test User"
username: "testuser"
role:     "Engineer"
theme:    "terminal"
`
	path := writeTempConfig(t, yml)

	cfg, err := config.Load(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Name != "Test User" {
		t.Errorf("Name = %q, want %q", cfg.Name, "Test User")
	}
	if cfg.Username != "testuser" {
		t.Errorf("Username = %q, want %q", cfg.Username, "testuser")
	}
	if cfg.Theme != "terminal" {
		t.Errorf("Theme = %q, want %q", cfg.Theme, "terminal")
	}
}

func TestLoad_DefaultTheme(t *testing.T) {
	yml := `
name:     "Test User"
username: "testuser"
role:     "Engineer"
`
	path := writeTempConfig(t, yml)

	cfg, err := config.Load(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Theme != "terminal" {
		t.Errorf("expected default theme %q, got %q", "terminal", cfg.Theme)
	}
}

func TestLoad_InvalidYAML(t *testing.T) {
	yml := `name: [invalid yaml`
	path := writeTempConfig(t, yml)

	_, err := config.Load(path)
	if err == nil {
		t.Error("expected error for invalid YAML, got nil")
	}
}

func TestLoad_MissingFile(t *testing.T) {
	_, err := config.Load("/no/such/file/config.yml")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}

func TestValidate_RequiredFields(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *config.Config
		wantErr bool
	}{
		{
			name:    "valid config",
			cfg:     &config.Config{Name: "Test", Username: "test", Role: "Engineer"},
			wantErr: false,
		},
		{
			name:    "missing name",
			cfg:     &config.Config{Username: "test", Role: "Engineer"},
			wantErr: true,
		},
		{
			name:    "missing username",
			cfg:     &config.Config{Name: "Test", Role: "Engineer"},
			wantErr: true,
		},
		{
			name:    "missing role",
			cfg:     &config.Config{Name: "Test", Username: "test"},
			wantErr: true,
		},
		{
			name:    "invalid website URL",
			cfg:     &config.Config{Name: "Test", Username: "test", Role: "Eng", Website: "not-a-url"},
			wantErr: true,
		},
		{
			name:    "valid website URL",
			cfg:     &config.Config{Name: "Test", Username: "test", Role: "Eng", Website: "https://example.com"},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			errs := config.Validate(tc.cfg)
			if tc.wantErr && !errs.HasErrors() {
				t.Error("expected validation errors, got none")
			}
			if !tc.wantErr && errs.HasErrors() {
				t.Errorf("unexpected validation errors: %v", errs)
			}
		})
	}
}

func TestValidate_ProjectsMissingName(t *testing.T) {
	cfg := &config.Config{
		Name:     "Test",
		Username: "test",
		Role:     "Engineer",
		Projects: []config.Project{
			{Name: "", Description: "desc", URL: "https://github.com/u/r"},
		},
	}

	errs := config.Validate(cfg)
	if !errs.HasErrors() {
		t.Error("expected error for project with empty name")
	}
}

// writeTempConfig creates a temporary config.yml for testing.
func writeTempConfig(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yml")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("writing temp config: %v", err)
	}
	return path
}
