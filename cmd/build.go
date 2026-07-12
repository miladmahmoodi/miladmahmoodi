package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/miladmahmoodi/miladmahmoodi/internal/generator"
)

var (
	buildConfig string
	buildOutput string
	buildDry    bool
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Generate README.md from config.yml",
	Long: `Runs the full generation pipeline:

  config.yml → theme → plugins → README.md

By default, profilegen build reads ./config.yml and writes ./README.md.`,
	Example: `  profilegen build
  profilegen build --config profile.yml --output README.md
  profilegen build --dry-run`,
	RunE: func(cmd *cobra.Command, args []string) error {
		start := time.Now()

		fmt.Println("  profilegen build")
		fmt.Printf("  config   %s\n", buildConfig)

		result, err := generator.Generate(generator.Options{
			ConfigPath: buildConfig,
			OutputPath: buildOutput,
			ThemesFS:   ThemesFS,
			DryRun:     buildDry,
		})
		if err != nil {
			return err
		}

		elapsed := time.Since(start)
		fmt.Printf("  theme    %s\n", result.ThemeName)
		fmt.Printf("  output   %s (%d bytes)\n", result.OutputPath, result.BytesOut)
		fmt.Printf("  done in  %dms\n", elapsed.Milliseconds())
		return nil
	},
}

func init() {
	buildCmd.Flags().StringVarP(&buildConfig, "config", "c", "config.yml", "path to config.yml")
	buildCmd.Flags().StringVarP(&buildOutput, "output", "o", "", "output path (default: README.md next to config)")
	buildCmd.Flags().BoolVar(&buildDry, "dry-run", false, "print output without writing to disk")
	rootCmd.AddCommand(buildCmd)
}
