package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	rx "github.com/pixel365/goreydenx"
	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

type creds struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,gte=6"`
}

type _email struct {
	Email string `validate:"required,email"`
}

type _password struct {
	Password string `validate:"required,gte=6"`
}

func NewLoginCommands() *cobra.Command {
	return &cobra.Command{
		Use:     "login",
		Aliases: []string{"l"},
		Short:   "Authorization by login and password",
		Long:    "Use this command to get a Token to use with the Reyden-X API",
		Run: func(cmd *cobra.Command, args []string) {
			config := helpers.LoadConfig()
			creds := &creds{}
			username(creds, config)
			password(creds)
			login(creds, config)
		},
	}
}

func username(creds *creds, config *helpers.Config) {
	_email := &_email{}
Loop:
	for {
		color.Cyan("Email:")
		reader := bufio.NewReader(os.Stdin)
		email, _ := reader.ReadString('\n')
		_email.Email = strings.Replace(email, "\n", "", -1)
		validate := validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(_email)
		if err != nil {
			color.Red("invalid email")
		} else {
			creds.Email = _email.Email
			if account := config.IsAccountExists(creds.Email); account != nil {
				if account.IsValid() {
					config.SetCurrent(account)
					helpers.SaveConfig(config)
					color.Green("Already authenticated")
					os.Exit(0)
				} else {
					password(creds)
					login(creds, config)
				}
			}
			break Loop
		}
	}
}

func password(creds *creds) {
	_password := &_password{}
Loop:
	for {
		color.Cyan("Password:")
		password, _ := term.ReadPassword(0)
		_password.Password = string(password)
		validate := validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(_password)
		if err != nil {
			color.Red("password length must be greater than equal 6 symbols")
		} else {
			creds.Password = _password.Password
			break Loop
		}
	}
}

func login(creds *creds, config *helpers.Config) {
	client := rx.NewClient(creds.Email, creds.Password).Auth()
	if client.Token != nil {
		token := helpers.Creds{
			Email:       creds.Email,
			AccessToken: client.Token.AccessToken,
			ExpiresIn:   client.Token.ExpiresIn,
			TokenType:   client.Token.TokenType,
		}
		config.SetCurrent(&token)
		helpers.SaveConfig(config)
		color.Green(helpers.Ok)
		os.Exit(0)
	} else {
		color.Red("invalid email or password")
	}
}
