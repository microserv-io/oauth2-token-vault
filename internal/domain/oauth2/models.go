package oauth2

import "time"

type Config struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
	AuthURL      string
	TokenURL     string
	RedirectURL  string
}

type Token struct {
	TokenType    string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}
