package main

import "github.com/urfave/cli/v2"

var statusCommand = &cli.Command{
	Name:  "status",
	Usage: "display the statuses of managed service(s)",
	Description: `
Display the status of the specified services. If no services are given, the
statuses of all services will be displayed.

This command does not change the state of the Noodle server. It will tell
you if the server is running or not, but it will not start it.
	`,
	ArgsUsage: "[service ...]",
	Action:    status,
}

func status(c *cli.Context) error {
	return nil
}
