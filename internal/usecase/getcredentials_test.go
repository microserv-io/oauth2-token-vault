package usecase

import (
	"context"
	"errors"
	"golang.org/x/oauth2"
	"testing"
	"time"

	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
)

type mockTokenSource struct {
	token *oauth2.Token
	err   error
}

func (m mockTokenSource) Token() (*oauth2.Token, error) {
	return m.token, m.err
}

type mockTokenSourceFactory struct {
	token *oauth2.Token
	err   error
}

func (m *mockTokenSourceFactory) NewTokenSource(_ context.Context, _ provider.Provider, _ oauthapp.OAuthApp) oauth2.TokenSource {
	return mockTokenSource{
		token: m.token,
		err:   m.err,
	}
}

type mockOAuthAppRepository struct {
	oauthapp.Repository
	oauthApps map[string]*oauthapp.OAuthApp
}

func (m *mockOAuthAppRepository) Find(_ context.Context, _ string, id string) (*oauthapp.OAuthApp, error) {
	if app, ok := m.oauthApps[id]; ok {
		return app, nil
	}
	return nil, errors.New("oauth app not found")
}

func (m *mockOAuthAppRepository) Update(_ context.Context, id string, updateFn func(app *oauthapp.OAuthApp) error) error {
	if app, ok := m.oauthApps[id]; ok {
		return updateFn(app)
	}
	return errors.New("oauth app not found")
}

type mockProviderRepository struct {
	provider.Repository
	providers map[string]*provider.Provider
}

func (m *mockProviderRepository) FindByName(_ context.Context, name string) (*provider.Provider, error) {
	if p, ok := m.providers[name]; ok {
		return p, nil
	}
	return nil, errors.New("provider not found")
}

func TestGetCredentialsUseCase_Execute(t *testing.T) {
	tests := []struct {
		name           string
		oauthApps      map[string]*oauthapp.OAuthApp
		providers      map[string]*provider.Provider
		token          *oauth2.Token
		tokenErr       error
		id             string
		ownerID        string
		wantCredential Credential
		wantErr        string
	}{
		{
			name: "valid credentials",
			oauthApps: map[string]*oauthapp.OAuthApp{
				"1": {
					ID:           "1",
					Provider:     "provider1",
					AccessToken:  "access_token",
					RefreshToken: "refresh_token",
					TokenType:    "Bearer",
					ExpiresAt:    time.Now().Add(-1 * time.Hour),
				},
			},
			providers: map[string]*provider.Provider{
				"provider1": {
					Name:         "provider1",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					Scopes:       []string{"scope1"},
					AuthURL:      "https://auth.url",
					TokenURL:     "https://token.url",
					RedirectURL:  "https://redirect.url",
				},
			},
			token: &oauth2.Token{
				AccessToken:  "new_access_token",
				RefreshToken: "new_refresh_token",
				Expiry:       time.Now().Add(1 * time.Hour),
				TokenType:    "Bearer",
			},
			id:      "1",
			ownerID: "owner1",
			wantCredential: Credential{
				TokenType:    "Bearer",
				ExpiresAt:    time.Now().Add(1 * time.Hour),
				AccessToken:  "new_access_token",
				RefreshToken: "new_refresh_token",
			},
			wantErr: "",
		},
		{
			name: "oauth app not found",
			oauthApps: map[string]*oauthapp.OAuthApp{
				"1": {
					ID:           "1",
					Provider:     "provider1",
					AccessToken:  "access_token",
					RefreshToken: "refresh_token",
					TokenType:    "Bearer",
					ExpiresAt:    time.Now().Add(-1 * time.Hour),
				},
			},
			providers: map[string]*provider.Provider{
				"provider1": {
					Name:         "provider1",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					Scopes:       []string{"scope1"},
					AuthURL:      "https://auth.url",
					TokenURL:     "https://token.url",
					RedirectURL:  "https://redirect.url",
				},
			},
			id:             "2",
			ownerID:        "owner1",
			wantCredential: Credential{},
			wantErr:        "oauth app not found",
		},
		{
			name: "provider not found",
			oauthApps: map[string]*oauthapp.OAuthApp{
				"1": {
					ID:           "1",
					Provider:     "provider1",
					AccessToken:  "access_token",
					RefreshToken: "refresh_token",
					TokenType:    "Bearer",
					ExpiresAt:    time.Now().Add(-1 * time.Hour),
				},
			},
			providers:      map[string]*provider.Provider{},
			id:             "1",
			ownerID:        "owner1",
			wantCredential: Credential{},
			wantErr:        "provider not found",
		},
		{
			name: "failed to refresh token",
			oauthApps: map[string]*oauthapp.OAuthApp{
				"1": {
					ID:           "1",
					Provider:     "provider1",
					AccessToken:  "access_token",
					RefreshToken: "refresh_token",
					TokenType:    "Bearer",
					ExpiresAt:    time.Now().Add(-1 * time.Hour),
				},
			},
			providers: map[string]*provider.Provider{
				"provider1": {
					Name:         "provider1",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					Scopes:       []string{"scope1"},
					AuthURL:      "https://auth.url",
					TokenURL:     "https://token.url",
					RedirectURL:  "https://redirect.url",
				},
			},
			token:          nil,
			tokenErr:       errors.New("failed to refresh token"),
			id:             "1",
			ownerID:        "owner1",
			wantCredential: Credential{},
			wantErr:        "failed to refresh token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oauthAppRepo := &mockOAuthAppRepository{oauthApps: tt.oauthApps}
			providerRepo := &mockProviderRepository{providers: tt.providers}
			tokenSourceFactory := &mockTokenSourceFactory{token: tt.token, err: tt.tokenErr}
			useCase := NewGetCredentialsUseCase(oauthAppRepo, providerRepo, tokenSourceFactory)

			got, err := useCase.Execute(context.Background(), tt.id, tt.ownerID)

			if tt.wantErr != "" {
				if err == nil || errors.Unwrap(err).Error() != tt.wantErr {
					t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Execute() error = %v", err)
				return
			}
			if got.AccessToken != tt.wantCredential.AccessToken {
				t.Errorf("Execute() got = %v, want %v", got.AccessToken, tt.wantCredential.AccessToken)
			}
			if got.RefreshToken != tt.wantCredential.RefreshToken {
				t.Errorf("Execute() got = %v, want %v", got.RefreshToken, tt.wantCredential.RefreshToken)
			}
			if got.TokenType != tt.wantCredential.TokenType {
				t.Errorf("Execute() got = %v, want %v", got.TokenType, tt.wantCredential.TokenType)
			}
		})
	}
}
