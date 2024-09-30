package tokensource

import (
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/url"
)

type Option func(*Factory) error

var _, _ = WithOAuthClient, WithEndpoint

// WithOAuthClient sets the OAuth client to use
func WithOAuthClient(client oauthcredentials.OAuthServiceClient) Option {
	return func(factory *Factory) error {
		factory.oauthClient = client
		return nil
	}
}

// WithEndpoint sets the endpoint to use
func WithEndpoint(endpoint *url.URL) Option {
	return func(factory *Factory) error {
		conn, err := grpc.NewClient(endpoint.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return fmt.Errorf("failed to connect to server: %w", err)
		}

		factory.oauthClient = oauthcredentials.NewOAuthServiceClient(conn)
		return nil
	}
}
