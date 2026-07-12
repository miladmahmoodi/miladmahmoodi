package cmd

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

const Version = "0.1.0"

// ThemesFS holds the embedded themes filesystem, injected from main.
var ThemesFS fs.FS

var rootCmd = &cobra.Command{
	Use:     "forge",
	Short:   "Build developer profiles, not README files.",
	Version: Version,
	Long: `
  ███████╗ ██████╗ ██████╗  ██████╗ ███████╗
  ██╔════╝██╔═══██╗██╔══██╗██╔════╝ ██╔════╝
  █████╗  ██║   ██║██████╔╝██║  ███╗█████╗
  ██╔══╝  ██║   ██║██╔══██╗██║   ██║██╔══╝
  ██║     ╚██████╔╝██║  ██║╚██████╔╝███████╗
  ╚═╝      ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚══════╝

  Build developer profiles, not README files.
  https://github.com/miladmahmoodi/miladmahmoodi`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute is the entry point called from main.go.
func Execute(embeddedFS fs.FS) {
	ThemesFS = embeddedFS

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "forge: "+err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate("forge v{{.Version}}\n")
}
