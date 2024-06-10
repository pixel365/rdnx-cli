package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewAccountsCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "accounts",
		Aliases: []string{"a"},
		Short:   "List of added accounts",
		Long:    "Use this command to view the entire list of added accounts and, if necessary, switch between them",
		Run: func(cmd *cobra.Command, args []string) {
			config := helpers.LoadConfig()
			total := len(config.Accounts)

			if total < 1 {
				fmt.Println("No accounts yet")
				os.Exit(0)
			}

			printAccounts(config)

			done := false
			for !done {
				choice := helpers.AskStringValue("Do you want to switch to another account? (y/n):", false)
				switch strings.ToLower(choice) {
				case "y":
					done = selectAccount(config)
				case "n":
					done = true
					os.Exit(0)
				default:
					color.Red("Invalid value: enter 'y' or 'n'")
				}
			}
		},
	}
}

func printAccounts(config *helpers.Config) {
	yellow := color.New(color.FgYellow).SprintfFunc()
	white := color.New(color.FgWhite).SprintfFunc()
	red := color.New(color.FgRed).SprintfFunc()
	green := color.New(color.FgGreen).SprintfFunc()
	i := 1

	for _, account := range config.Accounts {
		current := ""
		status := green("active")
		if !account.IsValid() {
			status = red("token expired")
		}

		if account.Email == config.CurrentAccountName() {
			current = " <- current"
		}

		s := fmt.Sprintf("%s %s (%s)%s", yellow(fmt.Sprintf("%d)", i)), white(account.Email), status, green(current))
		fmt.Println(s)
		i++
	}
}

func selectAccount(config *helpers.Config) bool {
	done := false
	for !done {
		value := helpers.AskIntegerValue("Enter your account number from the list above:", "Invalid account number", false)
		total := len(config.Accounts)
		if value > total {
			color.Red(fmt.Sprintf("Invalid choice. Value must be from 1 to %d", total))
		} else {
			i := 1
			for _, creds := range config.Accounts {
				if i == value {
					if creds.IsValid() {
						config.SetCurrent(&creds)
						helpers.SaveConfig(config)
						color.Green(helpers.Ok)
					} else {
						token := goreydenx.Token{
							AccessToken: creds.AccessToken,
							ExpiresIn:   creds.ExpiresIn,
							TokenType:   creds.TokenType,
						}
						client := goreydenx.NewClientWithToken(&token)
						newToken, err := client.RefreshToken()
						if err != nil {
							config.RemoveAccount(creds.Email)
							helpers.SaveConfig(config)
							color.Red(fmt.Sprintf("Unable to refresh token. Reason: %s", err.Error()))
							color.White(fmt.Sprintf("Account '%s' has been removed from the list. To receive a new token, re-authenticate using your login and password.", creds.Email))
							os.Exit(1)
						}

						creds.AccessToken = newToken.AccessToken
						creds.ExpiresIn = newToken.ExpiresIn
						creds.TokenType = newToken.TokenType

						config.SetCurrent(&creds)
						helpers.SaveConfig(config)
						color.Green("Successful login")
					}
					done = true
					break
				}
				i++
			}
		}
	}

	return done
}
