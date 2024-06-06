package traffic

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/pixel365/goreydenx/traffic"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewTrafficDevicesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     helpers.Devices,
		Aliases: []string{"d"},
		Short:   "Get information about the current volume of traffic by device",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client := config.RxClient()
			if client != nil {
				result, err := traffic.Devices(client)
				if err != nil {
					color.Red(err.Error())
					os.Exit(1)
				}

				j, _ := json.MarshalIndent(result, "", "    ")
				fmt.Println(string(j))
			} else {
				color.Red(helpers.Unauthorized)
			}
		},
	}
}
