package tokensource

import (
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
	"google.golang.org/grpc"
	"net/url"
)

type Option func(*Factory) error

func WithOAuthClient(client oauthcredentials.OAuthServiceClient) Option {
	return func(factory *Factory) error {
		factory.oauthClient = client
		return nil
	}
}

func WithEndpoint(endpoint *url.URL) Option {
	return func(factory *Factory) error {
		conn, err := grpc.NewClient(endpoint.Host)
		if err != nil {
			return fmt.Errorf("failed to connect to server: %w", err)
		}

		factory.oauthClient = oauthcredentials.NewOAuthServiceClient(conn)
		return nil
	}
}
