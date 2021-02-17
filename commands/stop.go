package main

import (
	"github.com/urfave/cli/v2"
)

var stopCommand = &cli.Command{
	Name:  "stop",
	Usage: "stop service(s)",
	Description: `
Stop the specified services. If no services are given, all started services
will be stopped.

The command waits for services to stop before exiting, so any services you
ask to be stopped will be stopped when the command exits without error, not
at some indeterminate time after that.

If this command stops the last running service, it will also stop the Noodle
server (and wait for it to stop, too).
	`,
	ArgsUsage: "[service ...]",
	Action:    stop,
}

func stop(c *cli.Context) error {
	return nil
}
