package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/miladmahmoodi/miladmahmoodi/internal/theme"
)

var themeCmd = &cobra.Command{
	Use:   "theme",
	Short: "Manage themes",
}

var themeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available themes",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("  forge theme list")
		fmt.Println()

		themes := theme.List(ThemesFS)
		for _, t := range themes {
			tag := ""
			if t.IsDefault {
				tag = "  (default)"
			}
			fmt.Printf("  %-14s  %s%s\n", t.Name, t.Description, tag)
		}

		fmt.Println()
		fmt.Println("  Set your theme in config.yml:")
		fmt.Println("    theme: terminal")
		return nil
	},
}

var themeInstallCmd = &cobra.Command{
	Use:   "install <name>",
	Short: "Install a remote theme",
	Long: `Downloads a theme from the Forge theme registry or a GitHub repository.

Examples:
  forge theme install minimal
  forge theme install github.com/user/forge-theme-retro`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Printf("  forge theme install %s\n\n", name)
		fmt.Println("  Remote theme installation is coming in v0.2.0.")
		fmt.Println("  To use a custom theme, place it in ./themes/<name>/")
		return nil
	},
}

func init() {
	themeCmd.AddCommand(themeListCmd)
	themeCmd.AddCommand(themeInstallCmd)
	rootCmd.AddCommand(themeCmd)
}
