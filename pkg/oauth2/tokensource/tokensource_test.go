package tokensource

import (
	"context"
	"testing"

	"github.com/microserv-io/oauth2-token-vault/pkg/proto/oauthcredentials/v1"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
)

type mockOAuthServiceClient struct {
	oauthcredentials.OAuthServiceClient
}

func (m *mockOAuthServiceClient) GetOAuthCredentialByProvider(_ context.Context, _ *oauthcredentials.GetOAuthCredentialByProviderRequest, _ ...grpc.CallOption) (*oauthcredentials.GetOAuthCredentialByProviderResponse, error) {
	return &oauthcredentials.GetOAuthCredentialByProviderResponse{
		AccessToken: "mockAccessToken",
	}, nil
}

func TestFactory_CreateTokenSource(t *testing.T) {
	tests := []struct {
		name          string
		opts          []Option
		provider      string
		resourceOwner string
		wantToken     *oauth2.Token
		wantErr       bool
	}{
		{
			name: "valid token source",
			opts: []Option{
				WithOAuthClient(&mockOAuthServiceClient{}),
			},
			provider:      "provider",
			resourceOwner: "resourceOwner",
			wantToken: &oauth2.Token{
				AccessToken: "mockAccessToken",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := NewFactory(tt.opts...)
			tokenSource := factory.CreateTokenSource(context.Background(), tt.provider, tt.resourceOwner)
			token, err := tokenSource.Token()
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenSource.Token() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if token.AccessToken != tt.wantToken.AccessToken {
				t.Errorf("TokenSource.Token() = %v, want %v", token.AccessToken, tt.wantToken.AccessToken)
			}
		})
	}
}
