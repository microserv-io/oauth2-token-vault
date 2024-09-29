package oauth2

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"net/url"
)

type TokenSourceFactory struct {
	oauthClient oauthcredentials.OAuthServiceClient
}

func NewTokenSourceFactory(oauthClient oauthcredentials.OAuthServiceClient) *TokenSourceFactory {
	return &TokenSourceFactory{oauthClient: oauthClient}
}

func NewStandardTokenSourceFactory(credentialsServerURL *url.URL) (*TokenSourceFactory, error) {
	conn, err := grpc.NewClient(credentialsServerURL.Host)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	oauthClient := oauthcredentials.NewOAuthServiceClient(conn)

	return &TokenSourceFactory{oauthClient: oauthClient}, nil
}

func (t TokenSourceFactory) CreateTokenSource(ctx context.Context, provider string, resourceOwner string) *TokenSource {
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
