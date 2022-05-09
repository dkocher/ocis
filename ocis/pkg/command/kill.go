package command

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/owncloud/ocis/v2/ocis-pkg/config"
	"github.com/owncloud/ocis/v2/ocis/pkg/register"
	"github.com/urfave/cli/v2"
)

// KillCommand is the entrypoint for the kill command.
func KillCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "kill",
		Usage:    "kill an extension by name in the runtime (supervised mode)",
		Category: "runtime",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "hostname",
				Value:       "localhost",
				EnvVars:     []string{"OCIS_RUNTIME_HOST"},
				Destination: &cfg.Runtime.Host,
			},
			&cli.StringFlag{
				Name:        "port",
				Value:       "9250",
				EnvVars:     []string{"OCIS_RUNTIME_PORT"},
				Destination: &cfg.Runtime.Port,
			},
		},
		Action: func(c *cli.Context) error {
			client, err := rpc.DialHTTP("tcp", net.JoinHostPort(cfg.Runtime.Host, cfg.Runtime.Port))
			if err != nil {
				log.Fatalf("Failed to connect to the runtime. Has the runtime been started and did you configure the right runtime address (\"%s\")", cfg.Runtime.Host+":"+cfg.Runtime.Port)
			}

			var arg1 int

			if err := client.Call("Service.Kill", os.Args[2], &arg1); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("process %v terminated", os.Args[2])

			return nil
		},
	}
}

func init() {
	register.AddCommand(KillCommand)
}