package command

import (
	"context"
	"fmt"
	"github.com/owncloud/ocis/v2/ocis-pkg/service/http"
	"github.com/owncloud/ocis/v2/ocis-pkg/version"
	"github.com/owncloud/ocis/v2/services/hub/pkg/service"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
)

// Server is the entrypoint for the server command.
func Server() *cli.Command {
	return &cli.Command{
		Name:     "server",
		Usage:    fmt.Sprintf("start the %s service without runtime (unsupervised mode)", "hub"),
		Category: "server",
		Action: func(c *cli.Context) error {
			var (
				ctx, cancel = func() (context.Context, context.CancelFunc) {
					return context.WithCancel(context.Background())
				}()
			)

			defer cancel()

			httpService, err := http.NewService(
				http.Name("hub"),
				http.Namespace("com.owncloud.web"),
				http.Version(version.GetString()),
				http.Address("127.0.0.1:9180"),
				http.Context(ctx),
			)
			if err != nil {
				return err
			}

			if err := micro.RegisterHandler(httpService.Server(), service.New()); err != nil {
				return err
			}

			return httpService.Run()
		},
	}
}
