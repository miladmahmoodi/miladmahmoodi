package plugin

// DefaultRegistry returns the built-in plugins that ship with Forge.
// Third-party plugins are registered by the user via config.yml and loaded
// from their own binaries (future: plugin protocol via stdin/stdout).
func DefaultRegistry() map[string]Plugin {
	return map[string]Plugin{
		// Built-in plugins will be added here as they are implemented.
		// The registry is intentionally empty in v0.1.0 to ship with zero
		// external API calls. Each plugin is opt-in.
	}
}

// BuiltinPluginInfo describes a plugin for listing purposes.
type BuiltinPluginInfo struct {
	Name        string
	Description string
	Status      string
}

// BuiltinPlugins is the catalog of first-party plugins.
var BuiltinPlugins = []BuiltinPluginInfo{
	{"github-activity", "Recent GitHub activity feed", "planned"},
	{"blog-rss", "Latest blog posts via RSS", "planned"},
	{"devto", "Latest dev.to articles", "planned"},
	{"hashnode", "Latest Hashnode articles", "planned"},
	{"spotify", "Currently listening on Spotify", "planned"},
	{"visitor-counter", "Profile visitor counter badge", "planned"},
	{"weather", "Current weather at your location", "planned"},
	{"quote", "Rotating developer quotes", "planned"},
}
