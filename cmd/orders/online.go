package orders

import (
	"github.com/fatih/color"
	"github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderOnlineCmd() (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "online",
		Aliases: []string{"o"},
		Short:   "Detailed information about users online",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client := config.RxClient()
			if client != nil {
				orderId := helpers.AskOrderId()
				result, err := orders.OnlineStats(client, uint32(orderId))
				helpers.Marshal(result, err)
			} else {
				color.Red(helpers.Unauthorized)
			}
		},
	}
	return
}
