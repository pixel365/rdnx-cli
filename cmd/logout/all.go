package logout

import (
	"github.com/fatih/color"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewLogoutAllCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "all",
		Aliases: []string{"a"},
		Short:   "Log out of all accounts",
		Long:    "Log out of all accounts and clear the configuration",
		Run: func(cmd *cobra.Command, args []string) {
			config := helpers.LoadConfig()
			if config.Logout(true) {
				helpers.SaveConfig(config)
				color.Green("Bye!")
			} else {
				color.Red("You are not logged in")
			}
		},
	}
}
