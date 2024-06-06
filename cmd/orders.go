package cmd

import (
	"github.com/fatih/color"
	m "github.com/pixel365/goreydenx/model"
	o "github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/cmd/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrdersCommands() (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "orders",
		Aliases: []string{"o"},
		Short:   "Order list",
		Long: `Use this command to get a list of all orders. 
				See also additional commands for obtaining other data on orders and operations on them`,
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client := config.RxClient()
			if client != nil {
				cursor := ""
				done := false
				for !done {
					result, err := o.Orders(client, cursor)
					helpers.Marshal(result, err)
					_next := helpers.Next[[]m.Order](result)
					cursor = result.Cursor
					done = !_next
				}
			} else {
				color.Red(helpers.Unauthorized)
			}
		},
	}
	command.AddCommand(orders.NewOrderAddViewsCmd())
	command.AddCommand(orders.NewOrderChangeOnlineCmd())
	command.AddCommand(orders.NewOrderClicksCmd())
	command.AddCommand(orders.NewOrderViewsCmd())
	command.AddCommand(orders.NewOrderOnlineCmd())
	command.AddCommand(orders.NewOrderSitesCmd())
	command.AddCommand(orders.NewOrderDetailsCmd())
	command.AddCommand(orders.NewOrderPaymentsCmd())
	command.AddCommand(orders.NewOrderMultiClicksCmd())
	command.AddCommand(orders.NewOrderMultiViewsCmd())
	command.AddCommand(orders.NewOrderRunCmd())
	command.AddCommand(orders.NewOrderStopCmd())
	command.AddCommand(orders.NewOrderCancelCmd())
	command.AddCommand(orders.NewOrderIncreaseChangeCmd())
	command.AddCommand(orders.NewOrderIncreaseOnCmd())
	command.AddCommand(orders.NewOrderIncreaseOffCmd())
	command.AddCommand(orders.NewCreateOrderCmd())
	return
}
