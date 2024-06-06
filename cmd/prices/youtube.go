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

func NewYouTubePricesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     helpers.YouTube,
		Aliases: []string{"yt"},
		Short:   "Get up-to-date information on available prices for new orders for YouTube",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client := config.RxClient()
			if client != nil {
				result, err := p.YouTube(client)
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