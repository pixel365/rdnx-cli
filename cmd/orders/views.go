package orders

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderViewsCommand(client **goreydenx.Client) (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "views",
		Aliases: []string{"v"},
		Short:   "Detailed information about views",
		Run: func(cmd *cobra.Command, args []string) {
			orderId := helpers.AskOrderId()
			result, err := orders.ViewsStats(*client, uint32(orderId))
			helpers.Marshal(result, err)
		},
	}
	return
}
