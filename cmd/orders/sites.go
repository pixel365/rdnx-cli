package orders

import (
	"github.com/fatih/color"
	"github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderSitesCmd() (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "sites",
		Aliases: []string{"s"},
		Short:   "Detailed information about sites",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client := config.RxClient()
			if client != nil {
				orderId := helpers.AskOrderId()
				result, err := orders.SitesStats(client, uint32(orderId))
				helpers.Marshal(result, err)
			} else {
				color.Red(helpers.Unauthorized)
			}
		},
	}
	return
}
