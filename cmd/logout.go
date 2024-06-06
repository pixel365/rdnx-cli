package cmd

import (
	"github.com/fatih/color"
	"github.com/pixel365/rdnx-cli/cmd/logout"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewLogoutCommands() (command *cobra.Command) {
	command = &cobra.Command{
		Use:     "logout",
		Aliases: []string{"exit", "bye"},
		Short:   "Log out of current account",
		Long:    "Log out of the current account, but save it in the configuration for later use",
		Run: func(cmd *cobra.Command, args []string) {
			config := helpers.LoadConfig()
			if config.Logout(false) {
				helpers.SaveConfig(config)
				color.Green("Bye!")
			} else {
				color.Red("You are not logged in")
			}
		},
	}
	command.AddCommand(logout.NewLogoutAllCmd())
	return
}
