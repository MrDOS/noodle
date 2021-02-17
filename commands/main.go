package main

import (
	_ "embed"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	// Embedded at build time with `-ldflags "-X 'main.description=...'"`.
	version string
	// The build process extracts the first paragraph of the README and drops
	// it into description.txt. This is easier than jumping through hoops to
	// escape any quotes in the text so that it can be safely passed through
	// the linker, then unescaped in code before being rendered.
	//go:embed description.txt
	description string

	app = &cli.App{
		Name:        "Noodle",
		HelpName:    "noodle",
		Usage:       "a project-oriented service manager",
		Version:     version,
		Description: description,
		Commands: []*cli.Command{
			statusCommand,
			startCommand,
			stopCommand,
			logsCommand,
			serverCommand,
		},
	}
)

func main() {
	app.Run(os.Args)
}
