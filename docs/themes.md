# Themes

profilegen ships six themes. Set your theme in `config.yml`:

```yaml
theme: terminal
```

## Available Themes

| Theme       | Description                          | Status     |
|-------------|--------------------------------------|------------|
| `terminal`  | Polished Linux terminal aesthetic    | ✓ stable   |
| `minimal`   | Clean, typography-focused layout     | ✓ stable   |
| `dashboard` | Card-based developer dashboard       | coming v0.2|
| `docs`      | Documentation-style layout           | coming v0.2|
| `retro`     | Retro ASCII art style                | coming v0.2|
| `hacker`    | Green-on-black matrix-inspired       | coming v0.2|

```bash
profilegen theme list
```

## Terminal Theme

The default theme renders your profile as a polished Linux terminal session.

- Animated SVG header (no GIF, no JS — GitHub-compatible)
- Each section renders as a `$` command with terminal output
- Easter egg section hidden in `<details>`
- Social badges via shields.io
- Featured projects with language and star count

## Minimal Theme

Clean, typography-first layout with no decorative elements.

- No badges
- Full-width content
- Simple bold headers for sections
- Suitable for designers and writers

## Building a Custom Theme

Place your theme in `./themes/<name>/`:

```
themes/
└── mytheme/
    ├── theme.yml
    └── templates/
        └── base.md.tmpl
```

**theme.yml:**

```yaml
name:        mytheme
description: My custom theme
author:      Your Name
version:     0.1.0
```

**templates/base.md.tmpl:**

```
# {{.Config.Name}}

{{.Config.Bio}}

{{range .Config.Skills}}
## {{.Category}}
{{range .Items}}- {{.}}
{{end}}
{{end}}
```

### Template Context

Every template receives a `RenderData` struct:

```go
type RenderData struct {
    Config      *config.Config  // parsed config.yml
    GeneratedAt string          // build timestamp
    Version     string          // profilegen version
}
```

### Built-in Template Functions

| Function | Description               | Example                         |
|----------|---------------------------|---------------------------------|
| `now`    | Current date string       | `{{now}}`                       |
| `add`    | Integer addition          | `{{add 1 2}}`                   |
| `join`   | Join strings with sep     | `{{join ", " .Items}}`          |
| `pad`    | Pad string to width       | `{{pad 12 .Category}}`          |

Standard Go template functions (`range`, `if`, `len`, `printf`, etc.) are all available.

## Publishing a Theme

To share your theme with the community, open a PR adding it to the `themes/` directory.

Theme registry and remote installation coming in v0.2.0.
