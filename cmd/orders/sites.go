package orders

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderSitesCommand(client **goreydenx.Client) (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "sites",
		Aliases: []string{"s"},
		Short:   "Detailed information about sites",
		Run: func(cmd *cobra.Command, args []string) {
			orderId := helpers.AskOrderId()
			result, err := orders.SitesStats(*client, uint32(orderId))
			helpers.Marshal(result, err)
		},
	}
	return
}
