package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewWhoAmICommand() *cobra.Command {
	return &cobra.Command{
		Use:     "whoami",
		Aliases: []string{"i"},
		Short:   "View account currently in use",
		PreRun: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
		},
		Run: func(cmd *cobra.Command, args []string) {
			config := helpers.LoadConfig()
			color.Green(fmt.Sprintf("You are logged in as \"%s\"", config.CurrentAccountName()))
		},
	}
}
