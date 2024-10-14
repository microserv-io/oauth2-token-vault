package oauth2

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/provider"
	"golang.org/x/oauth2"
)

func TestTokenSourceFactory_NewTokenSource(t *testing.T) {
	tests := []struct {
		name      string
		provider  provider.Provider
		oauthApp  oauthapp.OAuthApp
		wantToken *oauth2.Token
	}{
		{
			name: "valid token source",
			provider: provider.Provider{
				Name:         "provider1",
				ClientID:     "client_id",
				ClientSecret: "client_secret",
				Scopes:       []string{"scope1"},
				AuthURL:      "https://auth.url",
				TokenURL:     "https://token.url",
				RedirectURL:  "https://redirect.url",
			},
			oauthApp: oauthapp.OAuthApp{
				ID:           1,
				Provider:     "provider1",
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				TokenType:    "Bearer",
				ExpiresAt:    time.Now().Add(1 * time.Hour),
			},
			wantToken: &oauth2.Token{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				TokenType:    "Bearer",
				Expiry:       time.Now().Add(1 * time.Hour),
			},
		},
		{
			name: "expired token",
			provider: provider.Provider{
				Name:         "provider1",
				ClientID:     "client_id",
				ClientSecret: "client_secret",
				Scopes:       []string{"scope1"},
				AuthURL:      "https://auth.url",
				TokenURL:     "https://token.url",
				RedirectURL:  "https://redirect.url",
			},
			oauthApp: oauthapp.OAuthApp{
				ID:           1,
				Provider:     "provider1",
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				TokenType:    "Bearer",
				ExpiresAt:    time.Now().Add(-1 * time.Hour),
			},
			wantToken: &oauth2.Token{
				AccessToken:  "new_access_token",
				RefreshToken: "new_refresh_token",
				TokenType:    "Bearer",
				Expiry:       time.Now().Add(1 * time.Hour),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Create a mock HTTP server
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

			tt.provider.TokenURL = ts.URL

			factory := &TokenSourceFactory{}
			tokenSource := factory.NewTokenSource(context.Background(), &tt.provider, &tt.oauthApp)
			gotToken, err := tokenSource.Token()
			if err != nil {
				t.Fatalf("TokenSource.Token() error = %v", err)
			}
			if gotToken.AccessToken != tt.wantToken.AccessToken {
				t.Errorf("AccessToken = %v, want %v", gotToken.AccessToken, tt.wantToken.AccessToken)
			}
			if gotToken.RefreshToken != tt.wantToken.RefreshToken {
				t.Errorf("RefreshToken = %v, want %v", gotToken.RefreshToken, tt.wantToken.RefreshToken)
			}
			if gotToken.TokenType != tt.wantToken.TokenType {
				t.Errorf("TokenType = %v, want %v", gotToken.TokenType, tt.wantToken.TokenType)
			}
			if gotToken.Expiry.Sub(tt.wantToken.Expiry) > time.Second {
				t.Errorf("Expiry = %v, want %v", gotToken.Expiry, tt.wantToken.Expiry)
			}
		})
	}
}
