package cmd

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/spf13/cobra"

	"github.com/miladmahmoodi/miladmahmoodi/internal/generator"
)

var (
	publishConfig  string
	publishMessage string
	publishDryRun  bool
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Build README.md and push to GitHub",
	Long: `Runs forge build, then commits and pushes the generated README.md.

Requires git to be installed and the repository to have a remote configured.`,
	Example: `  forge publish
  forge publish --message "update profile"
  forge publish --dry-run`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("  forge publish")

		result, err := generator.Generate(generator.Options{
			ConfigPath: publishConfig,
			ThemesFS:   ThemesFS,
			DryRun:     publishDryRun,
		})
		if err != nil {
			return err
		}

		if publishDryRun {
			fmt.Println("  (dry-run) would commit and push README.md")
			return nil
		}

		fmt.Printf("  generated %s\n", result.OutputPath)

		msg := publishMessage
		if msg == "" {
			msg = fmt.Sprintf("chore: regenerate profile [forge %s] [skip ci]", time.Now().Format("2006-01-02"))
		}

		steps := [][]string{
			{"git", "add", result.OutputPath},
			{"git", "commit", "-m", msg},
			{"git", "push"},
		}

		for _, step := range steps {
			fmt.Printf("  $ %s\n", joinArgs(step))
			c := exec.Command(step[0], step[1:]...)
			c.Stdout = nil
			if out, err := c.CombinedOutput(); err != nil {
				return fmt.Errorf("git: %s\n%s", err, out)
			}
		}

		fmt.Println("  published")
		return nil
	},
}

func joinArgs(args []string) string {
	result := ""
	for i, a := range args {
		if i > 0 {
			result += " "
		}
		result += a
	}
	return result
}

func init() {
	publishCmd.Flags().StringVarP(&publishConfig, "config", "c", "config.yml", "path to config.yml")
	publishCmd.Flags().StringVarP(&publishMessage, "message", "m", "", "git commit message")
	publishCmd.Flags().BoolVar(&publishDryRun, "dry-run", false, "build but do not commit or push")
	rootCmd.AddCommand(publishCmd)
}
