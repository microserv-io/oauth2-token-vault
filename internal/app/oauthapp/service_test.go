package oauthapp

import (
	"context"
	"errors"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/app/oauthapp/mocks"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"net/url"
	"os"
	"testing"
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
			oauthAppRepository := new(mocks.MockOAuthAppRepository)
			providerRepository := new(mocks.MockProviderRepository)
			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

			service := NewOAuthAppService(oauthAppRepository, providerRepository, logger)

			oauthAppRepository.On("ListForOwner", mock.Anything, tt.ownerID).Return(tt.mockReturn, tt.mockError)

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
			oauthAppRepository := new(mocks.MockOAuthAppRepository)
			providerRepository := new(mocks.MockProviderRepository)
			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

			service := NewOAuthAppService(oauthAppRepository, providerRepository, logger)

			oauthAppRepository.On("Find", mock.Anything, tt.ownerID, tt.providerID).Return(tt.mockReturn, tt.mockError)

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
			oauthAppRepository := new(mocks.MockOAuthAppRepository)
			providerRepository := new(mocks.MockProviderRepository)
			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

			service := NewOAuthAppService(oauthAppRepository, providerRepository, logger)

			providerRepository.On("FindByName", mock.Anything, tt.providerID).Return(tt.mockProvider, tt.mockError)

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
