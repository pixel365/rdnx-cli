package orders

import (
	"github.com/fatih/color"
	"github.com/pixel365/goreydenx/action"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderAddViewsCmd() (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "add-views",
		Aliases: []string{"av"},
		Short:   "Add views to order",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client := config.RxClient()
			if client != nil {
				orderId := helpers.AskOrderId()
				value := helpers.AskIntegerValue(helpers.EnterIntValue, helpers.InvalidIntValue, false)
				result, err := action.AddViews(client, uint32(orderId), uint32(value))
				helpers.Marshal(result, err)
				helpers.WaitingTask(client, result)
			} else {
				color.Red(helpers.Unauthorized)
			}
		},
	}
	return
}
