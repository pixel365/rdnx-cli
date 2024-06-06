package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewWhoAmICommands() *cobra.Command {
	return &cobra.Command{
		Use:     "whoami",
		Aliases: []string{"i"},
		Short:   "View account currently in use",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			color.Green(fmt.Sprintf("You are logged in as \"%s\"", config.CurrentAccountName()))
		},
	}
}
