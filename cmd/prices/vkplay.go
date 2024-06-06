package prices

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	p "github.com/pixel365/goreydenx/prices"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewVkPlayPricesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     helpers.VkPLay,
		Aliases: []string{"vk"},
		Short:   "Get up-to-date information on available prices for new orders for VK Play",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client := config.RxClient()
			if client != nil {
				result, err := p.VkPlay(client)
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
