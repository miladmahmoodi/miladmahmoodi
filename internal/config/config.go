package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

// Config is the complete user profile configuration parsed from config.yml.
type Config struct {
	Name     string        `yaml:"name"`
	Username string        `yaml:"username"`
	Role     string        `yaml:"role"`
	Company  string        `yaml:"company"`
	Location string        `yaml:"location"`
	Website  string        `yaml:"website"`
	Bio      string        `yaml:"bio"`
	Socials  []Social      `yaml:"socials"`
	Skills   []SkillGroup  `yaml:"skills"`
	Projects []Project     `yaml:"projects"`
	Timeline []Event       `yaml:"timeline"`
	Articles []Article     `yaml:"articles"`
	Contact  Contact       `yaml:"contact"`
	Theme    string        `yaml:"theme"`
	Plugins  []PluginEntry `yaml:"plugins"`
}

type Social struct {
	Platform string `yaml:"platform"`
	URL      string `yaml:"url"`
	Handle   string `yaml:"handle"`
}

type SkillGroup struct {
	Category string   `yaml:"category"`
	Items    []string `yaml:"items"`
}

type Project struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	URL         string   `yaml:"url"`
	Stars       int      `yaml:"stars"`
	Language    string   `yaml:"language"`
	Tags        []string `yaml:"tags"`
	Featured    bool     `yaml:"featured"`
}

type Event struct {
	Year        string `yaml:"year"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

type Article struct {
	Title       string `yaml:"title"`
	URL         string `yaml:"url"`
	Platform    string `yaml:"platform"`
	PublishedAt string `yaml:"published_at"`
}

type Contact struct {
	Email   string `yaml:"email"`
	Twitter string `yaml:"twitter"`
	Discord string `yaml:"discord"`
}

type PluginEntry struct {
	Name    string         `yaml:"name"`
	Enabled bool           `yaml:"enabled"`
	Options map[string]any `yaml:"options"`
}

// Load reads and parses a config.yml file from the given path.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}

	cfg.applyDefaults()
	return &cfg, nil
}

func (c *Config) applyDefaults() {
	if c.Theme == "" {
		c.Theme = "terminal"
	}
}
