package main

import (
	"github.com/urfave/cli/v2"
)

var logsCommand = &cli.Command{
	Name:      "logs",
	Usage:     "retrieve/watch the logs for service(s)",
	ArgsUsage: "[service ...]",
	Action:    logs,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "f",
			Aliases: []string{"follow"},
			Usage:   "follow log output",
		},
		&cli.BoolFlag{
			Name:    "t",
			Aliases: []string{"timestamps"},
			Usage:   "prefix each log line with a timestamp",
		},
	},
}

func logs(c *cli.Context) error {
	return nil
}
