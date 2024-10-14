//go:build integration

package integration

import (
	"context"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/provider"
	"github.com/microserv-io/oauth2-token-vault/pkg/proto/oauthcredentials/v1"
	"github.com/microserv-io/oauth2-token-vault/tests"
	"google.golang.org/grpc"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOAuthByProvider(t *testing.T) {
	t.Parallel()

	beforeTestFunc := func(t *testing.T, app *tests.TestApp) {
		if err := app.Repositories.Provider.Create(context.Background(), &provider.Provider{
			Name: "Test Provider",
		}); err != nil {
			t.Fatalf("Failed to create provider: %v", err)
		}

		if err := app.Repositories.OAuthApp.Create(context.Background(), &oauthapp.OAuthApp{
			Provider: "Test Provider",
			OwnerID:  "Test Owner",
		}); err != nil {
		}
	}

	requestFunc := func(conn *grpc.ClientConn, input interface{}) (interface{}, error) {
		client := oauthcredentials.NewOAuthServiceClient(conn)
		return client.GetOAuthByProvider(context.Background(), input.(*oauthcredentials.GetOAuthByProviderRequest))
	}

	scenarios := []tests.GRPCScenario{
		{
			Name:           "success",
			BeforeTestFunc: beforeTestFunc,
			Input: &oauthcredentials.GetOAuthByProviderRequest{
				Provider: "Test Provider",
				Owner:    "Test Owner",
			},
			Request: requestFunc,
			ExpectedResp: &oauthcredentials.GetOAuthByProviderResponse{
				OauthApp: &oauthcredentials.OAuthApp{
					Id:       "1",
					Provider: "Test Provider",
					Owner:    "Test Owner",
				},
			},
		},
		{
			Name:           "no oauth app",
			BeforeTestFunc: beforeTestFunc,
			Input: &oauthcredentials.GetOAuthByProviderRequest{
				Provider: "Test Provider",
				Owner:    "Test Owner 2",
			},
			Request:     requestFunc,
			ExpectedErr: "record not found",
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGetOAuthCredentialByProvider(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.FormValue("refresh_token") != "test-refresh-token" {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error": "invalid_request"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"access_token": "test-access-token"}`))
	}))
	defer ts.Close()

	beforeTestFunc := func(t *testing.T, app *tests.TestApp) {
		if err := app.Repositories.Provider.Create(context.Background(), &provider.Provider{
			Name:     "Test Provider",
			TokenURL: ts.URL,
		}); err != nil {
			t.Fatalf("Failed to create provider: %v", err)
		}

		if err := app.Repositories.OAuthApp.Create(context.Background(), &oauthapp.OAuthApp{
			Provider:     "Test Provider",
			OwnerID:      "Test Owner",
			RefreshToken: "test-refresh-token",
		}); err != nil {
			t.Fatalf("Failed to create oauth app: %v", err)
		}

		if err := app.Repositories.OAuthApp.Create(context.Background(), &oauthapp.OAuthApp{
			Provider:     "Test Provider",
			OwnerID:      "Test Owner 2",
			RefreshToken: "invalid-refresh-token",
		}); err != nil {
			t.Fatalf("Failed to create oauth credential: %v", err)
		}
	}

	requestFunc := func(conn *grpc.ClientConn, input interface{}) (interface{}, error) {
		client := oauthcredentials.NewOAuthServiceClient(conn)
		return client.GetOAuthCredentialByProvider(context.Background(), input.(*oauthcredentials.GetOAuthCredentialByProviderRequest))
	}

	scenarios := []tests.GRPCScenario{
		{
			Name:           "success",
			BeforeTestFunc: beforeTestFunc,
			Input: &oauthcredentials.GetOAuthCredentialByProviderRequest{
				Provider: "Test Provider",
				Owner:    "Test Owner",
			},
			Request: requestFunc,
			ExpectedResp: &oauthcredentials.GetOAuthCredentialByProviderResponse{
				AccessToken: "test-access-token",
			},
		},
		{
			Name:           "error",
			BeforeTestFunc: beforeTestFunc,
			Input: &oauthcredentials.GetOAuthCredentialByProviderRequest{
				Provider: "Test Provider",
				Owner:    "Test Owner 2",
			},
			Request:     requestFunc,
			ExpectedErr: "invalid_request",
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
