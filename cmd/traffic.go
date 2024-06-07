package cmd

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/rdnx-cli/cmd/traffic"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewTrafficCommand() (command *cobra.Command) {
	var client *goreydenx.Client
	
	command = &cobra.Command{
		Use:     "traffic",
		Aliases: []string{"t"},
		Short:   "Get information about the current traffic volume",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client = config.RxClient()
		},
	}
	command.AddCommand(traffic.NewTrafficCountriesCommand(&client))
	command.AddCommand(traffic.NewTrafficLanguagesCommand(&client))
	command.AddCommand(traffic.NewTrafficDevicesCommand(&client))
	return
}
