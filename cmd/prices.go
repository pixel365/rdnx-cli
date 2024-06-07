package cmd

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/rdnx-cli/cmd/prices"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewPricesCommand() (command *cobra.Command) {
	var client *goreydenx.Client
	
	command = &cobra.Command{
		Use:     "prices",
		Aliases: []string{"p"},
		Short:   "Get up-to-date information on available prices",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client = config.RxClient()
		},
	}
	command.AddCommand(prices.NewTwitchPricesCommand(&client))
	command.AddCommand(prices.NewYouTubePricesCommand(&client))
	command.AddCommand(prices.NewTrovoPricesCommand(&client))
	command.AddCommand(prices.NewGoodGamePricesCommand(&client))
	command.AddCommand(prices.NewVkPlayPricesCommand(&client))
	return
}
