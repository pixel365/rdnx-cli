package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rdnx",
	Short: "Reyden-X CLI",
	Long: `Reyden-X is an automated service for promoting live broadcasts 
				on external sites with integrated system of viewers and views management.
				Complete documentation is available at https://api.reyden-x.com/docs`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(NewAccountsCommands())
	rootCmd.AddCommand(NewPricesCommands())
	rootCmd.AddCommand(NewMeCommands())
	rootCmd.AddCommand(NewWhoAmICommands())
	rootCmd.AddCommand(NewTrafficCommands())
	rootCmd.AddCommand(NewLogoutCommands())
	rootCmd.AddCommand(NewLoginCommands())
	rootCmd.AddCommand(NewOrdersCommands())
	rootCmd.AddCommand(NewInfoCommands())
	rootCmd.AddCommand(NewResetCommands())
}
