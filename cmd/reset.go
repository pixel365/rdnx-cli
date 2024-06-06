package cmd

import (
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewResetCommands() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Reset all settings",
		Long:  "Reset all settings and delete configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				choice := helpers.AskStringValue("Are you sure you want to delete the configuration file? (y/n)", false)
				switch strings.ToLower(choice) {
				case "y":
					path := helpers.GetConfigFilePath()
					if _, err := os.Stat(path); err == nil {
						err := os.Remove(path)
						if err == nil {
							color.Green("Done")
						}
					}
					os.Exit(0)
				case "n":
					os.Exit(0)
				default:
					color.Red("Invalid Value: Possible values only 'y' or 'n'")
				}
			}
		},
	}
}
