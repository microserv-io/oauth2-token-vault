package oauth2

import (
	"context"
	oauth "github.com/microserv-io/oauth2-token-vault/internal/domain/oauth2"
	"golang.org/x/oauth2"
)

var _ oauth.TokenSourceFactory = &TokenSourceFactory{}

type TokenSourceFactory struct {
}

func (t *TokenSourceFactory) NewTokenSource(
	ctx context.Context,
	tokenSourceConfig *oauth.TokenSourceConfig,
) (oauth2.TokenSource, error) {

	client := oauth2.Config{
		ClientID:     tokenSourceConfig.ClientID,
		ClientSecret: tokenSourceConfig.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  tokenSourceConfig.AuthURL,
			TokenURL: tokenSourceConfig.TokenURL,
		},
		RedirectURL: tokenSourceConfig.RedirectURL,
		Scopes:      tokenSourceConfig.Scopes,
	}

	return client.TokenSource(ctx, &oauth2.Token{
		AccessToken:  tokenSourceConfig.AccessToken,
		RefreshToken: tokenSourceConfig.RefreshToken,
		Expiry:       tokenSourceConfig.ExpiresAt,
		TokenType:    tokenSourceConfig.TokenType,
	}), nil
}
