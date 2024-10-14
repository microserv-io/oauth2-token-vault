package tokensource

import (
	"google.golang.org/grpc"
	"net/url"
	"testing"

	"github.com/microserv-io/oauth2-token-vault/pkg/proto/oauthcredentials/v1"
	"golang.org/x/net/context"
)

type mockOAuthServiceServer struct {
	oauthcredentials.OAuthServiceClient
}

func (m *mockOAuthServiceServer) GetOAuthCredentialByProvider(_ context.Context, _ *oauthcredentials.GetOAuthCredentialByProviderRequest, _ ...grpc.CallOption) (*oauthcredentials.GetOAuthCredentialByProviderResponse, error) {
	return &oauthcredentials.GetOAuthCredentialByProviderResponse{
		AccessToken: "mockAccessToken",
	}, nil
}

func TestWithOAuthClient(t *testing.T) {
	tests := []struct {
		name    string
		client  oauthcredentials.OAuthServiceClient
		wantErr bool
	}{
		{
			name:    "valid client",
			client:  &mockOAuthServiceServer{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := &Factory{}
			err := WithOAuthClient(tt.client)(factory)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithOAuthClient() error = %v, wantErr %v", err, tt.wantErr)
			}
			if factory.oauthClient != tt.client {
				t.Errorf("WithOAuthClient() = %v, want %v", factory.oauthClient, tt.client)
			}
		})
	}
}

func TestWithEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		endpoint *url.URL
		wantErr  bool
	}{
		{
			name:     "valid endpoint",
			endpoint: &url.URL{Host: "validhost"},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := &Factory{}

			if err := WithEndpoint(tt.endpoint)(factory); (err != nil) != tt.wantErr {
				t.Errorf("WithEndpoint() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && factory.oauthClient == nil {
				t.Errorf("WithEndpoint() = %v, want non-nil client", factory.oauthClient)
			}
		})
	}
}
