package cmd

import (
	"github.com/pixel365/goreydenx"
	o "github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/rdnx-cli/cmd/orders"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewOrdersCommand() (command *cobra.Command) {
	var client *goreydenx.Client

	command = &cobra.Command{
		Use:     "orders",
		Aliases: []string{"o"},
		Short:   "Order list",
		Long: `Use this command to get a list of all orders. 
				See also additional commands for obtaining other data on orders and operations on them`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client = config.RxClient()
		},
		Run: func(cmd *cobra.Command, args []string) {
			cursor := ""
			done := false
			for !done {
				result, err := o.Orders(client, cursor)
				helpers.Marshal(result, err)
				_next := helpers.Next(result)
				cursor = result.Cursor
				done = !_next
			}
		},
	}
	command.AddCommand(orders.NewOrderAddViewsCommand(&client))
	command.AddCommand(orders.NewOrderChangeOnlineCommand(&client))
	command.AddCommand(orders.NewOrderClicksCommand(&client))
	command.AddCommand(orders.NewOrderViewsCommand(&client))
	command.AddCommand(orders.NewOrderOnlineCommand(&client))
	command.AddCommand(orders.NewOrderSitesCommand(&client))
	command.AddCommand(orders.NewOrderDetailsCommand(&client))
	command.AddCommand(orders.NewOrderPaymentsCommand(&client))
	command.AddCommand(orders.NewOrderMultiClicksCommand(&client))
	command.AddCommand(orders.NewOrderMultiViewsCommand(&client))
	command.AddCommand(orders.NewOrderRunCommand(&client))
	command.AddCommand(orders.NewOrderStopCommand(&client))
	command.AddCommand(orders.NewOrderCancelCommand(&client))
	command.AddCommand(orders.NewOrderIncreaseChangeCommand(&client))
	command.AddCommand(orders.NewOrderIncreaseOnCommand(&client))
	command.AddCommand(orders.NewOrderIncreaseOffCommand(&client))
	command.AddCommand(orders.NewCreateOrderCommand(&client))
	command.AddCommand((orders.NewChangeLaunchModeCommand(&client)))
	return
}
