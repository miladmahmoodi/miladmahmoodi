package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var initForce bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Scaffold a new config.yml interactively",
	Long:  `Creates a starter config.yml by asking a few questions about you.`,
	Example: `  forge init
  forge init --force`,
	RunE: func(cmd *cobra.Command, args []string) error {
		const dest = "config.yml"

		if _, err := os.Stat(dest); err == nil && !initForce {
			return fmt.Errorf("%s already exists. Use --force to overwrite", dest)
		}

		fmt.Println("  forge init")
		fmt.Println()

		r := bufio.NewReader(os.Stdin)

		name := prompt(r, "  Your full name")
		username := prompt(r, "  GitHub username")
		role := prompt(r, "  Your role (e.g. Backend Engineer)")
		company := prompt(r, "  Company / org (optional)")
		location := prompt(r, "  Location (e.g. Tehran, Iran)")
		website := prompt(r, "  Website URL (optional)")
		twitter := prompt(r, "  Twitter handle (optional)")

		content := buildInitConfig(name, username, role, company, location, website, twitter)

		if err := os.WriteFile(dest, []byte(content), 0o644); err != nil {
			return fmt.Errorf("writing %s: %w", dest, err)
		}

		fmt.Println()
		fmt.Println("  Created config.yml")
		fmt.Println()
		fmt.Println("  Next steps:")
		fmt.Println("    1. Edit config.yml to add your projects, skills, and timeline")
		fmt.Println("    2. Run  forge build  to generate your README.md")
		fmt.Println("    3. Run  forge preview  to preview locally")
		return nil
	},
}

func prompt(r *bufio.Reader, label string) string {
	fmt.Printf("%s: ", label)
	line, _ := r.ReadString('\n')
	return strings.TrimSpace(line)
}

func buildInitConfig(name, username, role, company, location, website, twitter string) string {
	var sb strings.Builder

	sb.WriteString("# Forge config — https://github.com/miladmahmoodi/miladmahmoodi\n")
	sb.WriteString("# Edit this file, then run: forge build\n\n")

	sb.WriteString(fmt.Sprintf("name:     %q\n", name))
	sb.WriteString(fmt.Sprintf("username: %q\n", username))
	sb.WriteString(fmt.Sprintf("role:     %q\n", role))
	if company != "" {
		sb.WriteString(fmt.Sprintf("company:  %q\n", company))
	}
	if location != "" {
		sb.WriteString(fmt.Sprintf("location: %q\n", location))
	}
	if website != "" {
		sb.WriteString(fmt.Sprintf("website:  %q\n", website))
	}
	sb.WriteString(fmt.Sprintf("bio: %q\n", "I build things."))
	sb.WriteString("\ntheme: terminal\n")

	if twitter != "" {
		sb.WriteString("\nsocials:\n")
		sb.WriteString(fmt.Sprintf("  - platform: twitter\n    handle: %q\n    url: \"https://twitter.com/%s\"\n", twitter, strings.TrimPrefix(twitter, "@")))
		sb.WriteString("  - platform: github\n")
		sb.WriteString(fmt.Sprintf("    handle: %q\n", username))
		sb.WriteString(fmt.Sprintf("    url: \"https://github.com/%s\"\n", username))
	}

	sb.WriteString(`
skills:
  - category: Languages
    items: []
  - category: Frameworks
    items: []
  - category: Databases
    items: []
  - category: DevOps
    items: []

projects: []
  # - name: "My Project"
  #   description: "What it does"
  #   url: "https://github.com/user/repo"
  #   language: Go
  #   featured: true

timeline: []
  # - year: "2024"
  #   title: "Joined Acme Corp"
  #   description: "Started as a Backend Engineer"

contact:
  email: ""
`)

	return sb.String()
}

func init() {
	initCmd.Flags().BoolVar(&initForce, "force", false, "overwrite existing config.yml")
	rootCmd.AddCommand(initCmd)
}
