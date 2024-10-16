//go:build integration

package integration

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/provider"
	"github.com/microserv-io/oauth2-token-vault/pkg/proto/oauthcredentials/v1"
	"github.com/microserv-io/oauth2-token-vault/tests"
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

	beforeTestFunc := func(t *testing.T, app *tests.TestApp) {
		err := app.Repositories.Provider.Create(context.Background(), &provider.Provider{
			Name:         "Test Provider",
			AuthURL:      "https://example.com/auth",
			TokenURL:     "https://example.com/token",
			RedirectURL:  "https://example.com/redirect",
			Scopes:       []string{"scope1"},
			ClientID:     "client1",
			ClientSecret: "secret1",
		})

		if err != nil {
			t.Fatalf("Failed to create provider: %v", err)
		}
	}

	requestFunc := func(conn *grpc.ClientConn, input interface{}) (interface{}, error) {
		client := oauthcredentials.NewOAuthProviderServiceClient(conn)
		resp, err := client.ListProviders(context.Background(), input.(*oauthcredentials.ListProvidersRequest))
		return resp, err
	}

	scenarios := []tests.GRPCScenario{
		{
			Name:           "list providers",
			BeforeTestFunc: beforeTestFunc,
			Input:          &oauthcredentials.ListProvidersRequest{},
			Request:        requestFunc,
			ExpectedResp: &oauthcredentials.ListProvidersResponse{
				OauthProviders: []*oauthcredentials.OAuthProvider{
					{
						Name:        "Test Provider",
						AuthUrl:     "https://example.com/auth",
						TokenUrl:    "https://example.com/token",
						RedirectUri: "https://example.com/redirect",
						ClientId:    "client1",
						Scopes:      []string{"scope1"},
					},
				},
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestUpdateProvider(t *testing.T) {
	t.Parallel()

	beforeTestFunc := func(t *testing.T, app *tests.TestApp) {
		err := app.Repositories.Provider.Create(context.Background(), &provider.Provider{
			Name:     "Test Provider",
			AuthURL:  "https://example.com/auth",
			TokenURL: "https://example.com/token",
		})

		if err != nil {
			t.Fatalf("Failed to create provider: %v", err)
		}
	}

	requestFunc := func(conn *grpc.ClientConn, input interface{}) (interface{}, error) {
		client := oauthcredentials.NewOAuthProviderServiceClient(conn)
		resp, err := client.UpdateProvider(context.Background(), input.(*oauthcredentials.UpdateProviderRequest))
		return resp, err
	}

	scenarios := []tests.GRPCScenario{
		{
			Name:           "update provider",
			BeforeTestFunc: beforeTestFunc,
			Input: &oauthcredentials.UpdateProviderRequest{
				Name:         "Test Provider",
				AuthUrl:      "https://example.com/auth-new",
				TokenUrl:     "https://example.com/token-new",
				RedirectUri:  "https://example.com/redirect-new",
				ClientId:     "new-client1",
				ClientSecret: "new-secret1",
			},
			Request: requestFunc,
			ExpectedResp: &oauthcredentials.UpdateProviderResponse{
				OauthProvider: &oauthcredentials.OAuthProvider{
					Name:        "Test Provider",
					AuthUrl:     "https://example.com/auth-new",
					TokenUrl:    "https://example.com/token-new",
					RedirectUri: "https://example.com/redirect-new",
					ClientId:    "new-client1",
				},
			},
		},
		{
			Name:           "update provider that does not exist",
			BeforeTestFunc: beforeTestFunc,
			Input: &oauthcredentials.UpdateProviderRequest{
				Name:     "Nonexistent Provider",
				AuthUrl:  "https://example.com/auth2",
				TokenUrl: "https://example.com/token2",
			},
			Request:     requestFunc,
			ExpectedErr: "record not found",
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestExchangeAuthorizationCode(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.FormValue("code") == "test-code" {
			w.Write([]byte(`{"access_token": "test-access-token"}`))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "invalid_request"}`))
		}
	}))

	defer ts.Close()

	requestFunc := func(conn *grpc.ClientConn, input interface{}) (interface{}, error) {
		client := oauthcredentials.NewOAuthProviderServiceClient(conn)
		resp, err := client.ExchangeAuthorizationCode(context.Background(), input.(*oauthcredentials.ExchangeAuthorizationCodeRequest))
		return resp, err
	}

	beforeTestFunc := func(t *testing.T, app *tests.TestApp) {
		err := app.Repositories.Provider.Create(context.Background(), &provider.Provider{
			Name:         "Test Provider",
			AuthURL:      "https://example.com/auth",
			TokenURL:     ts.URL,
			RedirectURL:  "https://example.com/redirect",
			Scopes:       []string{"scope1"},
			ClientID:     "client1",
			ClientSecret: "secret1",
		})

		if err != nil {
			t.Fatalf("Failed to create provider: %v", err)
		}
	}

	scenarios := []tests.GRPCScenario{
		{
			Name:           "exchange authorization code",
			BeforeTestFunc: beforeTestFunc,
			Input: &oauthcredentials.ExchangeAuthorizationCodeRequest{
				Provider: "Test Provider",
				Code:     "test-code",
				Owner:    "test-owner",
			},
			Request:      requestFunc,
			ExpectedResp: &oauthcredentials.ExchangeAuthorizationCodeResponse{},
		},
		{
			Name:           "exchange authorization code with invalid code",
			BeforeTestFunc: beforeTestFunc,
			Input: &oauthcredentials.ExchangeAuthorizationCodeRequest{
				Provider: "Test Provider",
				Code:     "invalid-code",
				Owner:    "test-owner",
			},
			Request:     requestFunc,
			ExpectedErr: "invalid_request",
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestDeleteProvider(t *testing.T) {
	t.Parallel()

	requestFunc := func(conn *grpc.ClientConn, input interface{}) (interface{}, error) {
		client := oauthcredentials.NewOAuthProviderServiceClient(conn)
		resp, err := client.DeleteProvider(context.Background(), input.(*oauthcredentials.DeleteProviderRequest))
		return resp, err
	}

	beforeTestFunc := func(t *testing.T, app *tests.TestApp) {
		if err := app.Repositories.Provider.Create(context.Background(), &provider.Provider{
			Name: "Test Provider",
		}); err != nil {
			t.Fatalf("Failed to create provider 1: %v", err)
		}

		if err := app.Repositories.Provider.Create(context.Background(), &provider.Provider{
			Name: "Test Provider 2",
		}); err != nil {
			t.Fatalf("Failed to create provider 2: %v", err)
		}

		if err := app.Repositories.OAuthApp.Create(context.Background(), &oauthapp.OAuthApp{
			Provider: "Test Provider 2",
		}); err != nil {
			t.Fatalf("Failed to create oauth app: %v", err)
		}
	}

	scenarios := []tests.GRPCScenario{
		{
			Name:           "delete provider",
			BeforeTestFunc: beforeTestFunc,
			Input: &oauthcredentials.DeleteProviderRequest{
				Name: "Test Provider",
			},
			Request:      requestFunc,
			ExpectedResp: &oauthcredentials.DeleteProviderResponse{},
		},
		{
			Name:           "delete provider that does not exist",
			BeforeTestFunc: beforeTestFunc,
			Input: &oauthcredentials.DeleteProviderRequest{
				Name: "Nonexistent Provider",
			},
			Request:     requestFunc,
			ExpectedErr: "record not found",
		},
		{
			Name:           "delete provider with associated oauth apps",
			BeforeTestFunc: beforeTestFunc,
			Input: &oauthcredentials.DeleteProviderRequest{
				Name: "Test Provider 2",
			},
			Request:     requestFunc,
			ExpectedErr: "provider has associated oauth apps",
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
