package helpers

import (
	rx "github.com/pixel365/goreydenx"
)

const anonymous = "anonymous"

type Config struct {
	Accounts       map[string]Creds `json:"accounts"`
	CurrentAccount string           `json:"current_account"`
}

func (o *Config) CurrentAccountName() string {
	if _, ok := o.Accounts[o.CurrentAccount]; !ok {
		return anonymous
	}

	return o.CurrentAccount
}

func (o *Config) IsAuthenticated() bool {
	return o.CurrentAccountName() != anonymous
}

func (o *Config) IsAccountExists(email string) *Creds {
	if creds, ok := o.Accounts[email]; !ok {
		return nil
	} else {
		return &creds
	}
}

func (o *Config) SetCurrent(creds *Creds) *Config {
	if o.Accounts == nil {
		o.Accounts = make(map[string]Creds)
	}

	o.CurrentAccount = creds.Email
	o.Accounts[creds.Email] = *creds

	return o
}

func (o *Config) GetCurrent() (creds Creds, found bool) {
	creds, found = o.Accounts[o.CurrentAccount]
	return
}

func (o *Config) RemoveAccount(email string) {
	_, found := o.Accounts[email]
	if found {
		delete(o.Accounts, email)
	}

	if o.CurrentAccount == email {
		o.CurrentAccount = ""
	}
}

func (o *Config) Logout(all bool) bool {
	if o.IsAuthenticated() {
		o.CurrentAccount = ""
		if all {
			o.clear()
		}
		return true
	}
	if all {
		o.clear()
		return true
	}
	return false
}

func (o *Config) RxClient() *rx.Client {
	if !o.IsAuthenticated() {
		return nil
	}

	creds, _ := o.GetCurrent()
	token := rx.Token{
		AccessToken: creds.AccessToken,
		ExpiresIn:   creds.ExpiresIn,
		TokenType:   creds.TokenType,
	}

	return rx.NewClientWithToken(&token)
}

func (o *Config) clear() {
	o.CurrentAccount = ""
	for k := range o.Accounts {
		delete(o.Accounts, k)
	}
}
