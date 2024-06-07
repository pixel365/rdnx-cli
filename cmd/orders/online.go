package orders

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderOnlineCommand(client **goreydenx.Client) (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "online",
		Aliases: []string{"o"},
		Short:   "Detailed information about users online",
		Run: func(cmd *cobra.Command, args []string) {
			orderId := helpers.AskOrderId()
			result, err := orders.OnlineStats(*client, uint32(orderId))
			helpers.Marshal(result, err)
		},
	}
	return
}
