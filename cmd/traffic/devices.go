package traffic

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/traffic"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewTrafficDevicesCommand(client **goreydenx.Client) *cobra.Command {
	return &cobra.Command{
		Use:     helpers.Devices,
		Aliases: []string{"d"},
		Short:   "Get information about the current volume of traffic by device",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := traffic.Devices(*client)
			helpers.Marshal(result, err)
		},
	}
}
