package command

import (
	"github.com/owncloud/ocis/v2/ocis-pkg/config"
	"github.com/owncloud/ocis/v2/ocis/pkg/command/helper"
	"github.com/owncloud/ocis/v2/ocis/pkg/register"
	"github.com/owncloud/ocis/v2/services/hub/pkg/command"
	"github.com/urfave/cli/v2"
)

// HubCommand is the entrypoint for the web command.
func HubCommand(_ *config.Config) *cli.Command {
	return &cli.Command{
		Name:        "hub",
		Usage:       helper.SubcommandDescription("hub"),
		Category:    "services",
		Subcommands: command.GetCommands(),
	}
}

func init() {
	register.AddCommand(HubCommand)
}
