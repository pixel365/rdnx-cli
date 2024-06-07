package orders

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrderMultiViewsCommand(client **goreydenx.Client) (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "multi-views",
		Aliases: []string{"mv"},
		Short:   "View statistics for multiple orders",
		Run: func(cmd *cobra.Command, args []string) {
			identifiers := helpers.AskMultipleIntValues()
			result, err := orders.MultipleViewsStats(*client, identifiers)
			helpers.Marshal(result, err)
		},
	}
	return
}
