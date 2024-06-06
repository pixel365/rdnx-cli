package cmd

import (
	"fmt"

	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewInfoCommands() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "View general settings",
		Run: func(cmd *cobra.Command, args []string) {
			config := helpers.LoadConfig()
			fmt.Printf("Configuration File: %s\n", helpers.GetConfigFilePath())
			fmt.Printf("Total Accounts: %d\n", len(config.Accounts))
		},
	}
}
