package tokensource

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth2-token-vault/pkg/proto/oauthcredentials/v1"
	"golang.org/x/oauth2"
)

var _ oauth2.TokenSource = &TokenSource{}

type Factory struct {
	oauthClient oauthcredentials.OAuthServiceClient
}

func NewFactory(opts ...Option) *Factory {
	factory := &Factory{}

	for _, option := range opts {
		_ = option(factory)
	}

	return factory
}

func (t Factory) CreateTokenSource(ctx context.Context, provider string, resourceOwner string) *TokenSource {
	return NewTokenSource(ctx, t.oauthClient, provider, resourceOwner)
}

type TokenSource struct {
	context     context.Context
	oauthClient oauthcredentials.OAuthServiceClient

	provider      string
	resourceOwner string
}

func NewTokenSource(ctx context.Context, oauthClient oauthcredentials.OAuthServiceClient, provider string, resourceOwner string) *TokenSource {
	return &TokenSource{
		context:       ctx,
		oauthClient:   oauthClient,
		provider:      provider,
		resourceOwner: resourceOwner,
	}
}

func (t TokenSource) Token() (*oauth2.Token, error) {
	resp, err := t.oauthClient.GetOAuthCredentialByProvider(t.context, &oauthcredentials.GetOAuthCredentialByProviderRequest{
		Owner:    t.resourceOwner,
		Provider: t.provider,
	})

	if err != nil {
		return nil, fmt.Errorf("could not get oauth credentials: %w", err)
	}

	return &oauth2.Token{
		AccessToken: resp.GetAccessToken(),
	}, nil
}
