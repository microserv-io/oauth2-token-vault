package examples

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
	"google.golang.org/grpc"
)

func exchangeToken() {
	conn, err := grpc.NewClient("localhost:8080")
	if err != nil {
		panic(fmt.Errorf("failed to connect to server: %w", err))
	}

	client := oauthcredentials.NewOAuthServiceClient(conn)

	if _, err := client.ExchangeCodeForToken(context.TODO(), &oauthcredentials.ExchangeCodeForTokenRequest{
		Owner:    "some-user-id",
		Provider: "some-provider-id",
		Code:     "code",
		State:    "state-param",
	}); err != nil {
		panic(fmt.Errorf("failed to exchange code for token: %w", err))
	}
}
