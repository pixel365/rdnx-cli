package orders

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/action"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderStopCommand(client **goreydenx.Client) (command *cobra.Command) {
	command = &cobra.Command{
		Use:   "stop",
		Short: "Stop the order",
		Run: func(cmd *cobra.Command, args []string) {
			orderId := helpers.AskOrderId()
			result, err := action.Stop(*client, uint32(orderId))
			helpers.Marshal(result, err)
			helpers.WaitingTask(*client, result)
		},
	}
	return
}
