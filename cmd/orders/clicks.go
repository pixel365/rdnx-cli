package orders

import (
	"github.com/fatih/color"
	"github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderClicksCmd() (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "clicks",
		Aliases: []string{"c"},
		Short:   "Detailed information about clicks",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client := config.RxClient()
			if client != nil {
				orderId := helpers.AskOrderId()
				result, err := orders.ClicksStats(client, uint32(orderId))
				helpers.Marshal(result, err)
			} else {
				color.Red(helpers.Unauthorized)
			}
		},
	}
	return
}
