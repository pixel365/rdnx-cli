package traffic

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/traffic"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewTrafficCountriesCommand(client **goreydenx.Client) *cobra.Command {
	return &cobra.Command{
		Use:     helpers.Countries,
		Aliases: []string{"c"},
		Short:   "Get information about the current traffic volume by country",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := traffic.Countries(*client)
			helpers.Marshal(result, err)
		},
	}
}
