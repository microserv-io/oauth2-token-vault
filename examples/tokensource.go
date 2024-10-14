package examples

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth2-token-vault/pkg/oauth2/tokensource"
	"golang.org/x/oauth2"
	"net/url"
)

var _ = TokenSourceExample

// TokenSourceExample demonstrates how to use the token source to get a token
func TokenSourceExample() {

	serverEndpoint, err := url.Parse("oauth2-token-vault:8080")
	if err != nil {
		panic(fmt.Errorf("failed to parse server endpoint: %w", err))
	}

	tokenSourceFactory := tokensource.NewFactory(tokensource.WithEndpoint(serverEndpoint))

	// Use the token source to get a token
	httpClient := oauth2.NewClient(context.TODO(), tokenSourceFactory.CreateTokenSource(context.TODO(), "provider", "resource-owner"))

	// Use the http client to make requests
	_, _ = httpClient.Get("https://api.example.com")
}
