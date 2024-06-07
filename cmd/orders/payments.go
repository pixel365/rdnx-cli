package orders

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderPaymentsCommand(client **goreydenx.Client) (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "payments",
		Aliases: []string{"p"},
		Short:   "Order payments",
		Run: func(cmd *cobra.Command, args []string) {
			orderId := helpers.AskOrderId()
			cursor := ""
			done := false
			for !done {
				result, err := orders.Payments(*client, uint32(orderId), cursor)
				helpers.Marshal(result, err)
				_next := helpers.Next(result)
				cursor = result.Cursor
				done = !_next
			}
		},
	}
	return
}
