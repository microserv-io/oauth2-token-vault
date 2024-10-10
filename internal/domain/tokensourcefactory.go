package domain

import (
	"context"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	"golang.org/x/oauth2"
)

type TokenSourceFactory interface {
	NewTokenSource(ctx context.Context, provider *provider.Provider, oauthApp *oauthapp.OAuthApp) oauth2.TokenSource
}
