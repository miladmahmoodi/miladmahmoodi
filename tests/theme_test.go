package tests

import (
	"strings"
	"testing"
	"testing/fstest"

	"github.com/miladmahmoodi/miladmahmoodi/internal/config"
	"github.com/miladmahmoodi/miladmahmoodi/internal/theme"
)

func TestThemeLoad_EmbeddedTheme(t *testing.T) {
	memFS := fstest.MapFS{
		"themes/mytest/theme.yml": &fstest.MapFile{
			Data: []byte("name: mytest\ndescription: test theme\nauthor: test\nversion: 0.0.1\n"),
		},
		"themes/mytest/templates/base.md.tmpl": &fstest.MapFile{
			Data: []byte("Hello, {{.Config.Name}}!"),
		},
	}

	th, err := theme.Load("mytest", memFS)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if th.Name() != "mytest" {
		t.Errorf("Name() = %q, want %q", th.Name(), "mytest")
	}
}

func TestThemeLoad_MissingTheme(t *testing.T) {
	memFS := fstest.MapFS{}
	_, err := theme.Load("nosuchtheme", memFS)
	if err == nil {
		t.Error("expected error for missing theme, got nil")
	}
}

func TestThemeLoad_MissingThemeYML(t *testing.T) {
	memFS := fstest.MapFS{
		"themes/broken/templates/base.md.tmpl": &fstest.MapFile{
			Data: []byte("hello"),
		},
	}
	_, err := theme.Load("broken", memFS)
	if err == nil {
		t.Error("expected error for missing theme.yml, got nil")
	}
}

func TestThemeRender_TemplateContext(t *testing.T) {
	memFS := fstest.MapFS{
		"themes/ctx/theme.yml": &fstest.MapFile{
			Data: []byte("name: ctx\ndescription: context test\nauthor: test\nversion: 0.0.1\n"),
		},
		"themes/ctx/templates/base.md.tmpl": &fstest.MapFile{
			Data: []byte("Name:{{.Config.Name}} Role:{{.Config.Role}} Ver:{{.Version}}"),
		},
	}

	th, err := theme.Load("ctx", memFS)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	data := theme.NewRenderData(&config.Config{
		Name: "TestUser",
		Role: "Developer",
	}, "1.2.3")

	output, err := th.Render(data)
	if err != nil {
		t.Fatalf("Render() error: %v", err)
	}

	if !strings.Contains(output, "TestUser") {
		t.Error("output missing Name")
	}
	if !strings.Contains(output, "Developer") {
		t.Error("output missing Role")
	}
	if !strings.Contains(output, "1.2.3") {
		t.Error("output missing Version")
	}
}
