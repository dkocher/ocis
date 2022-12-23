package command

import (
	"context"
	ociscfg "github.com/owncloud/ocis/v2/ocis-pkg/config"
	"os"

	"github.com/owncloud/ocis/v2/ocis-pkg/clihelper"
	"github.com/thejerf/suture/v4"
	"github.com/urfave/cli/v2"
)

// GetCommands provides all commands for this service
func GetCommands() cli.Commands {
	return []*cli.Command{
		Server(),
	}
}

// Execute is the entry point for the web command.
func Execute() error {
	app := clihelper.DefaultApp(&cli.App{
		Name:     "hub",
		Usage:    "Serve ownCloud hub for oCIS",
		Commands: GetCommands(),
	})

	return app.Run(os.Args)
}

// SutureService allows for the web command to be embedded and supervised by a suture supervisor tree.
type SutureService struct{}

// NewSutureService creates a new web.SutureService
func NewSutureService(_ *ociscfg.Config) suture.Service {
	return SutureService{}
}

func (s SutureService) Serve(_ context.Context) error {
	if err := Execute(); err != nil {
		return err
	}

	return nil
}
