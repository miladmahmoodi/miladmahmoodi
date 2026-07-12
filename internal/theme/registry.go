package theme

import (
	"io/fs"
	"sort"
)

// BuiltinThemes lists the themes shipped with Forge.
var BuiltinThemes = []ThemeInfo{
	{Name: "terminal", Description: "Polished Linux terminal aesthetic", IsDefault: true},
	{Name: "minimal", Description: "Clean, minimal typography-focused layout"},
	{Name: "dashboard", Description: "Card-based developer dashboard"},
	{Name: "docs", Description: "Documentation-style layout"},
	{Name: "retro", Description: "Retro ASCII art style"},
	{Name: "hacker", Description: "Green-on-black matrix-inspired"},
}

// ThemeInfo is a lightweight descriptor used for listing.
type ThemeInfo struct {
	Name        string
	Description string
	IsDefault   bool
}

// List returns all themes discoverable from embedded and local directories.
func List(embedded fs.FS) []ThemeInfo {
	infos := make([]ThemeInfo, len(BuiltinThemes))
	copy(infos, BuiltinThemes)

	entries, err := fs.ReadDir(embedded, "themes")
	if err != nil {
		return infos
	}

	known := map[string]bool{}
	for _, t := range BuiltinThemes {
		known[t.Name] = true
	}

	for _, e := range entries {
		if e.IsDir() && !known[e.Name()] {
			infos = append(infos, ThemeInfo{Name: e.Name(), Description: "custom"})
		}
	}

	sort.Slice(infos, func(i, j int) bool {
		if infos[i].IsDefault {
			return true
		}
		return infos[i].Name < infos[j].Name
	})

	return infos
}
