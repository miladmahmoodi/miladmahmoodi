# Plugins

Plugins extend Forge with dynamic content fetched at build time.

## Using Plugins

Enable plugins in `config.yml`:

```yaml
plugins:
  - name:    github-activity
    enabled: true
    options:
      limit: 5

  - name:    blog-rss
    enabled: true
    options:
      url: "https://yourblog.dev/rss"
      limit: 3
```

Disabled plugins are skipped entirely:

```yaml
plugins:
  - name:    spotify
    enabled: false   # won't run
```

## Available Plugins

```bash
forge plugin list
```

| Plugin            | Description                        | Status       |
|-------------------|------------------------------------|--------------|
| `github-activity` | Recent GitHub activity feed        | planned v0.2 |
| `blog-rss`        | Latest blog posts via RSS          | planned v0.2 |
| `devto`           | Latest dev.to articles             | planned v0.2 |
| `hashnode`        | Latest Hashnode articles           | planned v0.2 |
| `spotify`         | Currently listening on Spotify     | planned v0.2 |
| `visitor-counter` | Profile visitor counter badge      | planned v0.2 |
| `weather`         | Current weather at your location   | planned v0.2 |
| `quote`           | Rotating developer quotes          | planned v0.2 |

## Building a Plugin

Implement the `Plugin` interface:

```go
package myplugin

import "github.com/miladmahmoodi/forge/internal/config"

type MyPlugin struct{}

func (p *MyPlugin) Name() string { return "my-plugin" }

func (p *MyPlugin) Render(cfg *config.Config, opts map[string]any) (string, error) {
    // fetch data, build markdown, return it
    return "### My Plugin Output\n\nHello from my plugin.", nil
}
```

Register it in your theme or via a Forge plugin extension point (v0.2.0).

## Plugin Output

Plugin output is injected into the template at the plugin's designated slot.  
Plugins return a markdown string — they do not have access to the template engine.

## Design Principles

- **Opt-in only** — no plugin runs unless explicitly enabled
- **No side effects** — plugins only produce output strings
- **Fail gracefully** — a failing plugin logs a warning, never crashes the build
- **Zero network calls by default** — the core build never phones home
