package main

import (
	"github.com/urfave/cli/v2"
)

var startCommand = &cli.Command{
	Name:  "start",
	Usage: "start service(s)",
	Description: `
Start the specified services. If no services are given, all stopped services
will be started.

The command waits for services to be started before exiting. However,
because Noodle doesn't have any concept of a “health check” for managed
services, you may have to wait some further amount of time while the
services go through their startup process before they can accept requests.

If the Noodle server isn't already running, this command will start it.
	`,
	ArgsUsage: "[service ...]",
	Action:    start,
}

func start(c *cli.Context) error {
	return nil
}
