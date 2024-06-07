package traffic

import (
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/traffic"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewTrafficLanguagesCommand(client **goreydenx.Client) *cobra.Command {
	return &cobra.Command{
		Use:     helpers.Languages,
		Aliases: []string{"l"},
		Short:   "Get information about the current volume of traffic by language",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := traffic.Languages(*client)
			helpers.Marshal(result, err)
		},
	}
}
