package oauth2

import (
	"context"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/provider"
	oauth "github.com/microserv-io/oauth2-token-vault/internal/domain/oauth2"
	"golang.org/x/oauth2"
)

var _ oauth.TokenSourceFactory = &TokenSourceFactory{}

type TokenSourceFactory struct {
}

func (t *TokenSourceFactory) NewTokenSource(ctx context.Context, provider *provider.Provider, oauthApp *oauthapp.OAuthApp) oauth2.TokenSource {
	client := oauth2.Config{
		ClientID:     provider.ClientID,
		ClientSecret: provider.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  provider.AuthURL,
			TokenURL: provider.TokenURL,
		},
		RedirectURL: provider.RedirectURL,
		Scopes:      provider.Scopes,
	}

	return client.TokenSource(ctx, &oauth2.Token{
		AccessToken:  oauthApp.AccessToken,
		RefreshToken: oauthApp.RefreshToken,
		Expiry:       oauthApp.ExpiresAt,
		TokenType:    oauthApp.TokenType,
	})
}
