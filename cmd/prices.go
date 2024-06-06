package cmd

import (
	"github.com/pixel365/rdnx-cli/cmd/prices"
	"github.com/spf13/cobra"
)

func NewPricesCommands() (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "prices",
		Aliases: []string{"p"},
		Short:   "Get up-to-date information on available prices",
	}
	command.AddCommand(prices.NewTwitchPricesCmd())
	command.AddCommand(prices.NewYouTubePricesCmd())
	command.AddCommand(prices.NewTrovoPricesCmd())
	command.AddCommand(prices.NewGoodGamePricesCmd())
	command.AddCommand(prices.NewVkPlayPricesCmd())
	return
}
