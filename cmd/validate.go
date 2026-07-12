package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/miladmahmoodi/forge/internal/config"
)

var validateConfig string

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate config.yml against the schema",
	Long:  `Checks that your config.yml is valid before running forge build.`,
	Example: `  forge validate
  forge validate --config profile.yml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("  forge validate  %s\n\n", validateConfig)

		cfg, err := config.Load(validateConfig)
		if err != nil {
			return fmt.Errorf("cannot parse config: %w", err)
		}

		errs := config.Validate(cfg)
		if errs.HasErrors() {
			fmt.Println("  errors found:")
			fmt.Println(errs.Error())
			fmt.Println()
			return fmt.Errorf("config is invalid")
		}

		fmt.Println("  config.yml is valid")
		fmt.Printf("  theme:    %s\n", cfg.Theme)
		fmt.Printf("  name:     %s\n", cfg.Name)
		fmt.Printf("  projects: %d\n", len(cfg.Projects))
		fmt.Printf("  timeline: %d events\n", len(cfg.Timeline))
		fmt.Printf("  plugins:  %d\n", len(cfg.Plugins))
		return nil
	},
}

func init() {
	validateCmd.Flags().StringVarP(&validateConfig, "config", "c", "config.yml", "path to config.yml")
	rootCmd.AddCommand(validateCmd)
}
