package cmd

import (
	"github.com/pixel365/rdnx-cli/cmd/traffic"
	"github.com/spf13/cobra"
)

func NewTrafficCommands() (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "traffic",
		Aliases: []string{"t"},
		Short:   "Get information about the current traffic volume",
	}
	command.AddCommand(traffic.NewTrafficCountriesCmd())
	command.AddCommand(traffic.NewTrafficLanguagesCmd())
	command.AddCommand(traffic.NewTrafficDevicesCmd())
	return
}
