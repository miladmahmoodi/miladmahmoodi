package main

import (
	"embed"

	"github.com/miladmahmoodi/forge/cmd"
)

//go:embed themes
var themesFS embed.FS

func main() {
	cmd.Execute(themesFS)
}
