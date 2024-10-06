package oauthapp

import (
	"context"
	"errors"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
	"log/slog"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestOAuthAppService_ListOAuths(t *testing.T) {
	tests := []struct {
		name          string
		ownerID       string
		mockReturn    []*oauthapp.OAuthApp
		mockError     error
		expectedError bool
	}{
		{
			name:          "Success",
			ownerID:       "owner1",
			mockReturn:    []*oauthapp.OAuthApp{{ID: "app1"}},
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Failure",
			ownerID:       "owner2",
			mockReturn:    nil,
			mockError:     errors.New("database error"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oauthAppRepository := NewMockOAuthAppRepository(t)
			providerRepository := NewMockProviderRepository(t)
			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

			service := NewService(oauthAppRepository, providerRepository, nil, logger)

			oauthAppRepository.EXPECT().ListForOwner(mock.Anything, tt.ownerID).Return(tt.mockReturn, tt.mockError)

			result, err := service.ListOAuthsForOwner(context.Background(), tt.ownerID)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Len(t, result, 1)
			}
		})
	}
}

func TestOAuthAppService_GetOAuthByID(t *testing.T) {
	tests := []struct {
		name          string
		providerID    string
		ownerID       string
		mockReturn    *oauthapp.OAuthApp
		mockError     error
		expectedError bool
	}{
		{
			name:          "Success",
			providerID:    "provider1",
			ownerID:       "owner1",
			mockReturn:    &oauthapp.OAuthApp{ID: "app1"},
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Failure",
			providerID:    "provider2",
			ownerID:       "owner2",
			mockReturn:    nil,
			mockError:     errors.New("database error"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oauthAppRepository := NewMockOAuthAppRepository(t)
			providerRepository := NewMockProviderRepository(t)
			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

			service := NewService(oauthAppRepository, providerRepository, nil, logger)

			oauthAppRepository.EXPECT().Find(mock.Anything, tt.ownerID, tt.providerID).Return(tt.mockReturn, tt.mockError)

			result, err := service.GetOAuthForProviderAndOwner(context.Background(), tt.providerID, tt.ownerID)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestOAuthAppService_CreateAuthorizationURL(t *testing.T) {
	tests := []struct {
		name          string
		providerID    string
		scopes        []string
		state         string
		mockProvider  *provider.Provider
		mockError     error
		expectedError bool
	}{
		{
			name:       "Success",
			providerID: "provider1",
			scopes:     []string{"scope1"},
			state:      "state1",
			mockProvider: &provider.Provider{
				ClientID:     "client1",
				ClientSecret: "secret1",
				RedirectURL:  "http://localhost",
				AuthURL:      "http://auth",
				TokenURL:     "http://token",
				Scopes:       []string{"scope1"},
			},
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Failure",
			providerID:    "provider2",
			scopes:        []string{"scope2"},
			state:         "state2",
			mockProvider:  nil,
			mockError:     errors.New("provider not found"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oauthAppRepository := NewMockOAuthAppRepository(t)
			providerRepository := NewMockProviderRepository(t)

			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

			service := NewService(oauthAppRepository, providerRepository, nil, logger)

			providerRepository.EXPECT().FindByName(mock.Anything, tt.providerID).Return(tt.mockProvider, tt.mockError)

			result, err := service.CreateAuthorizationURLForProvider(context.Background(), tt.providerID, tt.scopes, tt.state)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s", tt.mockProvider.AuthURL, tt.mockProvider.ClientID, url.QueryEscape(tt.mockProvider.RedirectURL), "scope1", "state1"), result.URL.String())
			}
		})
	}
}

func TestService_RetrieveAccessToken(t *testing.T) {
	tests := []struct {
		name              string
		providerID        string
		ownerID           string
		mockOAuthApp      *oauthapp.OAuthApp
		mockProvider      *provider.Provider
		mockToken         *oauth2.Token
		mockFindError     error
		mockProviderError error
		mockTokenError    error
		expectedError     bool
	}{
		{
			name:       "Success",
			providerID: "provider1",
			ownerID:    "owner1",
			mockOAuthApp: &oauthapp.OAuthApp{
				ID:           "app1",
				Provider:     "provider1",
				AccessToken:  "oldAccessToken",
				RefreshToken: "oldRefreshToken",
				TokenType:    "Bearer",
				ExpiresAt:    time.Now().Add(-time.Hour),
			},
			mockProvider: &provider.Provider{
				ClientID:     "client1",
				ClientSecret: "secret1",
				TokenURL:     "http://token",
			},
			mockToken: &oauth2.Token{
				AccessToken:  "newAccessToken",
				RefreshToken: "newRefreshToken",
				TokenType:    "Bearer",
				Expiry:       time.Now().Add(time.Hour),
			},
			mockFindError:     nil,
			mockProviderError: nil,
			mockTokenError:    nil,
			expectedError:     false,
		},
		{
			name:              "Failure - Find Error",
			providerID:        "provider2",
			ownerID:           "owner2",
			mockOAuthApp:      nil,
			mockProvider:      nil,
			mockToken:         nil,
			mockFindError:     errors.New("find error"),
			mockProviderError: nil,
			mockTokenError:    nil,
			expectedError:     true,
		},
		{
			name:              "Failure - Provider Error",
			providerID:        "provider3",
			ownerID:           "owner3",
			mockOAuthApp:      &oauthapp.OAuthApp{},
			mockProvider:      nil,
			mockToken:         nil,
			mockFindError:     nil,
			mockProviderError: errors.New("provider error"),
			mockTokenError:    nil,
			expectedError:     true,
		},
		{
			name:       "Failure - Token Error",
			providerID: "provider4",
			ownerID:    "owner4",
			mockOAuthApp: &oauthapp.OAuthApp{
				ID:           "app4",
				Provider:     "provider4",
				AccessToken:  "oldAccessToken",
				RefreshToken: "oldRefreshToken",
				TokenType:    "Bearer",
				ExpiresAt:    time.Now().Add(-time.Hour),
			},
			mockProvider: &provider.Provider{
				ClientID:     "client4",
				ClientSecret: "secret4",
				TokenURL:     "http://token",
			},
			mockToken:         nil,
			mockFindError:     nil,
			mockProviderError: nil,
			mockTokenError:    errors.New("token error"),
			expectedError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oauthAppRepository := NewMockOAuthAppRepository(t)
			providerRepository := NewMockProviderRepository(t)
			tokenSourceFactory := NewMockTokenSourceFactory(t)

			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

			service := NewService(oauthAppRepository, providerRepository, tokenSourceFactory, logger)

			oauthAppRepository.On("Find", mock.Anything, tt.ownerID, tt.providerID).Return(tt.mockOAuthApp, tt.mockFindError)

			if tt.mockOAuthApp != nil {
				providerRepository.On("FindByName", mock.Anything, tt.providerID).Return(tt.mockProvider, tt.mockProviderError)
			}

			if tt.mockOAuthApp != nil && tt.mockProvider != nil {
				tokenSourceFactory.EXPECT().NewTokenSource(mock.Anything, *tt.mockProvider, *tt.mockOAuthApp).Return(mockTokenSource{token: tt.mockToken, err: tt.mockTokenError})
				if tt.mockToken != nil {
					oauthAppRepository.On("UpdateByID", mock.Anything, tt.mockOAuthApp.ID, mock.Anything).Return(nil)
				}
			}

			result, err := service.RetrieveAccessToken(context.Background(), tt.providerID, tt.ownerID)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.mockToken.AccessToken, result.AccessToken)
			}
		})
	}
}

type mockTokenSource struct {
	token *oauth2.Token
	err   error
}

func (m mockTokenSource) Token() (*oauth2.Token, error) {
	return m.token, m.err
}
