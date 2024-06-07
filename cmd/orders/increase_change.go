package orders

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/action"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderIncreaseChangeCommand(client **goreydenx.Client) (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "increase-change",
		Aliases: []string{"ic"},
		Short:   "Change the time of the smooth set of viewers",
		Run: func(cmd *cobra.Command, args []string) {
			orderId := helpers.AskOrderId()
			value := helpers.AskIntegerValue(helpers.EnterIntValue, helpers.InvalidIntValue, false)
			result, err := action.ChangeIncreaseValue(*client, uint32(orderId), uint32(value))
			helpers.Marshal(result, err)
			helpers.WaitingTask(*client, result)
		},
	}
	return
}
