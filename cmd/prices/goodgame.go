package prices

import (
	"github.com/pixel365/goreydenx"
	p "github.com/pixel365/goreydenx/prices"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewGoodGamePricesCommand(client **goreydenx.Client) *cobra.Command {
	return &cobra.Command{
		Use:     helpers.GoodGame,
		Aliases: []string{"g"},
		Short:   "Get up-to-date information on available prices for new orders for GoodGame",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := p.GoodGame(*client)
			helpers.Marshal(result, err)
		},
	}
}
