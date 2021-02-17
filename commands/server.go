package main

import (
	"github.com/urfave/cli/v2"
)

var serverCommand = &cli.Command{
	Name:  "server",
	Usage: "starts the Unix socket command listener",
	Description: `
Noodle will launch its own server when you start any managed service, and
will stop the server after the last service is stopped. Under normal
circumstances, you don't need to start the server yourself.

If you do manually start the server with this command, it will run in the
foreground without daemonizing/detaching itself.
	`,
	Action: server,
}

func server(c *cli.Context) error {
	return nil
}
