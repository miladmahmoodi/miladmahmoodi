package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check your environment for common issues",
	Long:  `Inspects your environment and reports anything that might affect forge.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("  forge doctor")
	fmt.Println()

		checks := []check{
			{"config.yml exists", checkFileExists("config.yml")},
			{"git available", checkBinary("git")},
			{"go available", checkBinary("go")},
			{"themes directory", checkThemesDir()},
		}

		allOk := true
		for _, c := range checks {
			status := "✓"
			if c.err != nil {
				status = "✗"
				allOk = false
			}
			if c.err != nil {
				fmt.Printf("  %s  %s\n    → %v\n", status, c.label, c.err)
			} else {
				fmt.Printf("  %s  %s\n", status, c.label)
			}
		}

		fmt.Println()
		fmt.Printf("  os/arch   %s/%s\n", runtime.GOOS, runtime.GOARCH)
		fmt.Printf("  forge     v%s\n", Version)

		if !allOk {
			fmt.Println()
			return fmt.Errorf("doctor found issues — see above")
		}

		fmt.Println()
		fmt.Println("  everything looks good")
		return nil
	},
}

type check struct {
	label string
	err   error
}

func checkFileExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%s not found — run: forge init", path)
	}
	return nil
}

func checkBinary(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		return fmt.Errorf("%s not found in PATH", name)
	}
	return nil
}

func checkThemesDir() error {
	if _, err := os.Stat("themes"); os.IsNotExist(err) {
		return nil // fine — using embedded themes
	}
	return nil
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
