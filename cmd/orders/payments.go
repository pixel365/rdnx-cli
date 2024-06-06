package orders

import (
	"github.com/fatih/color"
	"github.com/pixel365/goreydenx/model"
	"github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderPaymentsCmd() (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "payments",
		Aliases: []string{"p"},
		Short:   "Order payments",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client := config.RxClient()
			if client != nil {
				orderId := helpers.AskOrderId()
				cursor := ""
				done := false
				for !done {
					result, err := orders.Payments(client, uint32(orderId), cursor)
					helpers.Marshal(result, err)
					_next := helpers.Next[[]model.Payment](result)
					cursor = result.Cursor
					done = !_next
				}
			} else {
				color.Red(helpers.Unauthorized)
			}
		},
	}
	return
}
