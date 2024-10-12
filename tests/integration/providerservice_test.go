//go:build integration

package integration

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
	"github.com/microserv-io/oauth-credentials-server/tests"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO(roccolangeweg): Currently the ordering of the tests matter. Create must be executed first, delete must be executed last. Tests need to be standalone (e.g. unique IDs)

func TestCreateProvider(t *testing.T) {
	t.Parallel()
	app := tests.NewTestApp("config")

	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", app.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := oauthcredentials.NewOAuthProviderServiceClient(conn)

	// Create a new provider
	req := &oauthcredentials.CreateProviderRequest{
		Name:     "Test Provider",
		AuthUrl:  "https://example.com/auth",
		TokenUrl: "https://example.com/token",
		// Add other necessary fields here
	}

	resp, err := client.CreateProvider(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateProvider failed: %v", err)
	}

	assert.NotNil(t, resp)
	assert.Equal(t, "Test Provider", resp.GetOauthProvider().GetName())
	assert.Equal(t, "https://example.com/auth", resp.GetOauthProvider().GetAuthUrl())
	assert.Equal(t, "https://example.com/token", resp.GetOauthProvider().GetTokenUrl())

	app.Cleanup()
}

func TestListProviders(t *testing.T) {
	t.Parallel()

	app := tests.NewTestApp("config")

	err := app.Repositories.Provider.Create(context.Background(), &provider.Provider{
		Name:     "Test Provider",
		AuthURL:  "https://example.com/auth",
		TokenURL: "https://example.com/token",
	})

	// Set up the gRPC connection
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", app.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := oauthcredentials.NewOAuthProviderServiceClient(conn)

	// List all providers
	req := &oauthcredentials.ListProvidersRequest{}

	stream, err := client.ListProviders(context.Background(), req)
	if err != nil {
		t.Fatalf("ListProviders failed: %v", err)
	}

	var providers []*oauthcredentials.OAuthProvider
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		for _, p := range resp.GetOauthProviders() {
			providers = append(providers, p)
		}
	}

	assert.NotEmpty(t, providers)

	app.Cleanup()
}

func TestUpdateProvider(t *testing.T) {
	t.Parallel()

	app := tests.NewTestApp("config")

	err := app.Repositories.Provider.Create(context.Background(), &provider.Provider{
		Name:     "Test Provider",
		AuthURL:  "https://example.com/auth",
		TokenURL: "https://example.com/token",
	})

	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", app.Port), grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := oauthcredentials.NewOAuthProviderServiceClient(conn)

	// Update a provider
	req := &oauthcredentials.UpdateProviderRequest{
		Name:     "Test Provider",
		AuthUrl:  "https://example.com/auth2",
		TokenUrl: "https://example.com/token2",
		// Add other necessary fields here
	}

	resp, err := client.UpdateProvider(context.Background(), req)
	if err != nil {
		t.Fatalf("UpdateProvider failed: %v", err)
	}

	assert.NotNil(t, resp)
	assert.Equal(t, "Test Provider", resp.GetOauthProvider().GetName())
	assert.Equal(t, "https://example.com/auth2", resp.GetOauthProvider().GetAuthUrl())
	assert.Equal(t, "https://example.com/token2", resp.GetOauthProvider().GetTokenUrl())

	app.Cleanup()
}

func TestExchangeAuthorizationCode(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{
					"access_token": "new_access_token",
					"token_type": "Bearer",
					"refresh_token": "new_refresh_token",
					"expires_in": 3600
				}`)); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))

	defer ts.Close()

	app := tests.NewTestApp("config")

	app.Repositories.Provider.Create(context.Background(), &provider.Provider{
		Name:     "Test Provider",
		AuthURL:  "https://example.com/auth",
		TokenURL: ts.URL,
	})
	// Set up the gRPC connection
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", app.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := oauthcredentials.NewOAuthProviderServiceClient(conn)

	// Exchange an authorization code
	req := &oauthcredentials.ExchangeAuthorizationCodeRequest{
		Provider: "Test Provider",
		Code:     "test-code",
		Owner:    "test-owner",
	}

	resp, err := client.ExchangeAuthorizationCode(context.Background(), req)
	if err != nil {
		t.Fatalf("ExchangeAuthorizationCode failed: %v", err)
	}

	assert.NotNil(t, resp)

	app.Cleanup()
}

func TestDeleteProvider(t *testing.T) {
	// Set up the gRPC connection
	app := tests.NewTestApp("config")

	app.Repositories.Provider.Create(context.Background(), &provider.Provider{
		Name:     "Test Provider",
		AuthURL:  "https://example.com/auth",
		TokenURL: "https://example.com/token",
	})

	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", app.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := oauthcredentials.NewOAuthProviderServiceClient(conn)

	// Delete a provider
	req := &oauthcredentials.DeleteProviderRequest{
		Id: "Test Provider",
	}

	resp, err := client.DeleteProvider(context.Background(), req)
	if err != nil {
		t.Fatalf("DeleteProvider failed: %v", err)
	}

	assert.NotNil(t, resp)

	p, err := app.Repositories.Provider.FindByName(context.Background(), "Test Provider")
	assert.Nil(t, p)
	assert.NotNil(t, err)

	app.Cleanup()
}
