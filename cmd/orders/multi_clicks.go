package orders

import (
	"github.com/fatih/color"
	"github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderMultiClicksCmd() (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "multi-clicks",
		Aliases: []string{"mc"},
		Short:   "Click-through statistics for multiple orders",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client := config.RxClient()
			if client != nil {
				identifiers := helpers.AskMultipleIntValues()
				result, err := orders.MultipleClicksStats(client, identifiers)
				helpers.Marshal(result, err)
			} else {
				color.Red(helpers.Unauthorized)
			}
		},
	}
	return
}
