package orders

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderClicksCommand(client **goreydenx.Client) (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "clicks",
		Aliases: []string{"c"},
		Short:   "Detailed information about clicks",
		Run: func(cmd *cobra.Command, args []string) {
			orderId := helpers.AskOrderId()
			result, err := orders.ClicksStats(*client, uint32(orderId))
			helpers.Marshal(result, err)
		},
	}
	return
}
