package prices

import (
	"github.com/pixel365/goreydenx"
	p "github.com/pixel365/goreydenx/prices"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewYouTubePricesCommand(client **goreydenx.Client) *cobra.Command {
	return &cobra.Command{
		Use:     helpers.YouTube,
		Aliases: []string{"yt"},
		Short:   "Get up-to-date information on available prices for new orders for YouTube",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := p.YouTube(*client)
			helpers.Marshal(result, err)
		},
	}
}
