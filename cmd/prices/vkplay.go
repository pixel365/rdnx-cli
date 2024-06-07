package prices

import (
	"github.com/pixel365/goreydenx"
	p "github.com/pixel365/goreydenx/prices"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewVkPlayPricesCommand(client **goreydenx.Client) *cobra.Command {
	return &cobra.Command{
		Use:     helpers.VkPLay,
		Aliases: []string{"vk"},
		Short:   "Get up-to-date information on available prices for new orders for VK Play",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := p.VkPlay(*client)
			helpers.Marshal(result, err)
		},
	}
}
