package oauth2

import (
	"context"
	"golang.org/x/oauth2"
	"time"
)

type TokenSourceConfig struct {
	ClientID     string
	ClientSecret string
	AuthURL      string
	TokenURL     string
	RedirectURL  string
	Scopes       []string
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresAt    time.Time
}

type TokenSourceFactory interface {
	NewTokenSource(
		ctx context.Context,
		tokenSourceConfig *TokenSourceConfig,
	) (oauth2.TokenSource, error)
}
