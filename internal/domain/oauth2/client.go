package oauth2

import (
	"context"
	"time"
)

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

type Client interface {
	GetAuthorizationURL(config *Config, state string) (string, error)
	Exchange(ctx context.Context, config *Config, code string) (*Token, error)
}
