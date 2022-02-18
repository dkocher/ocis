package command

import (
	"fmt"

	"github.com/asim/go-micro/plugins/events/nats/v4"
	"github.com/cs3org/reva/pkg/events"
	"github.com/cs3org/reva/pkg/events/server"
	"github.com/owncloud/ocis/notifications/pkg/config"
	"github.com/owncloud/ocis/notifications/pkg/config/parser"
	"github.com/owncloud/ocis/notifications/pkg/logging"
	"github.com/owncloud/ocis/notifications/pkg/service"
	"github.com/urfave/cli/v2"
)

// Server is the entrypoint for the server command.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "server",
		Usage:    fmt.Sprintf("start %s extension without runtime (unsupervised mode)", cfg.Service.Name),
		Category: "server",
		Before: func(c *cli.Context) error {
			return parser.ParseConfig(cfg)
		},
		Action: func(c *cli.Context) error {
			logger := logging.Configure(cfg.Service.Name, cfg.Log)

			group := "notifications"

			evs := []events.Unmarshaller{
				events.ShareCreated{},
			}

			client, err := server.NewNatsStream(nats.Address("127.0.0.1:4222"), nats.ClusterID("test-cluster"))
			if err != nil {
				return err
			}
			evts, err := events.Consume(client, group, evs...)
			if err != nil {
				return err
			}

			svc := service.NewEventsNotifier(evts, logger)
			return svc.Run()
		},
	}
}
