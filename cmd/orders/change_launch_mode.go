package orders

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/action"
	m "github.com/pixel365/goreydenx/model"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewChangeLaunchModeCommand(client **goreydenx.Client) (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "change-launch-mode",
		Aliases: []string{"clm"},
		Short:   "Change order launch mode",
		Run: func(cmd *cobra.Command, args []string) {
			orderId := helpers.AskOrderId()
			launchMode, delayTime := helpers.AskLaunchMode()
			result, err := action.ChangeLaunchMode(*client, uint32(orderId), &m.LaunchParameters{
				LaunchMode: launchMode,
				DelayTime:  uint32(delayTime),
			})
			helpers.Marshal(result, err)
			helpers.WaitingTask(*client, result)
		},
	}
	return
}
