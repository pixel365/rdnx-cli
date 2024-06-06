package orders

import (
	"github.com/fatih/color"
	"github.com/pixel365/goreydenx/action"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderCancelCmd() (command *cobra.Command) {
	command = &cobra.Command{
		Use:   "cancel",
		Short: "Cancel the order",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client := config.RxClient()
			if client != nil {
				orderId := helpers.AskOrderId()
				result, err := action.Cancel(client, uint32(orderId))
				helpers.Marshal(result, err)
				helpers.WaitingTask(client, result)
			} else {
				color.Red(helpers.Unauthorized)
			}
		},
	}
	return
}
