package helpers

import "github.com/golang-module/carbon"

type Creds struct {
	Email       string `json:"email,omitempty"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type,omitempty"`
}

func (o *Creds) IsValid() bool {
	if o.AccessToken == "" || o.ExpiresIn == "" {
		return false
	}

	now := carbon.Now(carbon.UTC)
	return carbon.Parse(o.ExpiresIn).Compare(">=", now)
}
