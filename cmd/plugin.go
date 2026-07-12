package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/miladmahmoodi/miladmahmoodi/internal/plugin"
)

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Manage plugins",
}

var pluginListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available plugins",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("  profilegen plugin list")
		fmt.Println()

		for _, p := range plugin.BuiltinPlugins {
			fmt.Printf("  %-20s  %-40s  [%s]\n", p.Name, p.Description, p.Status)
		}

		fmt.Println()
		fmt.Println("  Enable plugins in config.yml:")
		fmt.Println("    plugins:")
		fmt.Println("      - name: github-activity")
		fmt.Println("        enabled: true")
		return nil
	},
}

var pluginInstallCmd = &cobra.Command{
	Use:   "install <name>",
	Short: "Install a community plugin",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Printf("  profilegen plugin install %s\n\n", name)
		fmt.Println("  Community plugin installation is coming in v0.2.0.")
		return nil
	},
}

func init() {
	pluginCmd.AddCommand(pluginListCmd)
	pluginCmd.AddCommand(pluginInstallCmd)
	rootCmd.AddCommand(pluginCmd)
}
