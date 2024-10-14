package oauth2

import (
	"context"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/provider"
	"golang.org/x/oauth2"
)

type TokenSourceFactory interface {
	NewTokenSource(ctx context.Context, provider *provider.Provider, oauthApp *oauthapp.OAuthApp) oauth2.TokenSource
}
