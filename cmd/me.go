package cmd

import (
	"os"
	"sync"

	"github.com/fatih/color"
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/user"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewMeCommand() *cobra.Command {
	var client *goreydenx.Client

	return &cobra.Command{
		Use:   "me",
		Short: "Up-to-date information about your current account and balance status",
		PreRun: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client = config.RxClient()
		},
		Run: func(cmd *cobra.Command, args []string) {
			var wg sync.WaitGroup
			wg.Add(2)

			var errors []error
			result := make(map[string]any)

			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				user, err := user.Account(client)
				if err != nil {
					errors = append(errors, err)
				} else {
					result["account"] = user
				}
			}(&wg)

			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				balance, err := user.Balance(client)
				if err != nil {
					errors = append(errors, err)
				} else {
					result["balance"] = balance
				}
			}(&wg)

			wg.Wait()

			if len(errors) > 0 {
				for _, e := range errors {
					color.Red(e.Error())
				}
				os.Exit(1)
			}

			helpers.Marshal(result, nil)
		},
	}
}
