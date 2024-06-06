package orders

import (
	"github.com/pixel365/rdnx-cli/cmd/orders/create"
	"github.com/spf13/cobra"
)

func NewCreateOrderCmd() (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "create",
		Aliases: []string{"new"},
		Short:   "Create new order",
	}
	command.AddCommand(create.NewTwitchOrderCmd())
	command.AddCommand(create.NewYouTubeOrderCmd())
	return
}
