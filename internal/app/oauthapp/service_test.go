package oauthapp

import (
	"context"
	"errors"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
	"log/slog"
	"os"
	"testing"
	"time"
)

func TestOAuthAppService_ListOAuthAppsForOwner(t *testing.T) {
	tests := []struct {
		name          string
		ownerID       string
		mockSetup     func(oauthAppRepository *MockOAuthAppRepository)
		expectedError bool
		expectedResp  *ListOAuthAppsForOwnerResponse
	}{
		{
			name:    "Success",
			ownerID: "owner1",
			mockSetup: func(oauthAppRepository *MockOAuthAppRepository) {
				oauthAppRepository.EXPECT().ListForOwner(mock.Anything, "owner1").Return([]*oauthapp.OAuthApp{
					{
						ID: 1,
					},
				}, nil)
			},
			expectedError: false,
			expectedResp: &ListOAuthAppsForOwnerResponse{
				Apps: []*OAuthApp{
					{
						ID: "1",
					},
				},
			},
		},
		{
			name:    "Failure",
			ownerID: "owner2",
			mockSetup: func(oauthAppRepository *MockOAuthAppRepository) {
				oauthAppRepository.EXPECT().ListForOwner(mock.Anything, "owner2").Return(nil, errors.New("database error"))
			},
			expectedError: true,
			expectedResp:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oauthAppRepository := NewMockOAuthAppRepository(t)

			tt.mockSetup(oauthAppRepository)

			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

			service := NewService(oauthAppRepository, nil, nil, logger)

			resp, err := service.ListOAuthAppsForOwner(context.Background(), tt.ownerID)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedResp, resp)
		})
	}
}

func TestOAuthAppService_GetOAuthForProviderAndOwner(t *testing.T) {
	tests := []struct {
		name          string
		providerID    string
		ownerID       string
		mockSetup     func(oauthAppRepository *MockOAuthAppRepository)
		expectedError bool
		expectedResp  *GetOAuthForProviderAndOwnerResponse
	}{
		{
			name:       "Success",
			providerID: "provider1",
			ownerID:    "owner1",
			mockSetup: func(oauthAppRepository *MockOAuthAppRepository) {
				oauthAppRepository.EXPECT().Find(mock.Anything, "owner1", "provider1").Return(&oauthapp.OAuthApp{
					ID: 1,
				}, nil)
			},
			expectedError: false,
			expectedResp: &GetOAuthForProviderAndOwnerResponse{
				App: &OAuthApp{
					ID: "1",
				},
			},
		},
		{
			name:       "Failure",
			providerID: "provider2",
			ownerID:    "owner2",
			mockSetup: func(oauthAppRepository *MockOAuthAppRepository) {
				oauthAppRepository.EXPECT().Find(mock.Anything, "owner2", "provider2").Return(nil, errors.New("database error"))
			},
			expectedError: true,
			expectedResp:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oauthAppRepository := NewMockOAuthAppRepository(t)

			tt.mockSetup(oauthAppRepository)

			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

			service := NewService(oauthAppRepository, nil, nil, logger)

			resp, err := service.GetOAuthForProviderAndOwner(context.Background(), tt.providerID, tt.ownerID)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedResp, resp)
		})
	}
}

func TestService_RetrieveAccessToken(t *testing.T) {
	tests := []struct {
		name          string
		providerID    string
		ownerID       string
		mockSetup     func(oauthAppRepository *MockOAuthAppRepository, providerRepository *MockProviderRepository, tokenSourceFactory *MockTokenSourceFactory)
		expectedError bool
		expectedResp  *RetrieveAccessTokenResponse
	}{
		{
			name:       "Success",
			providerID: "provider1",
			ownerID:    "owner1",
			mockSetup: func(oauthAppRepository *MockOAuthAppRepository, providerRepository *MockProviderRepository, tokenSourceFactory *MockTokenSourceFactory) {

				mockOAuthApp := &oauthapp.OAuthApp{
					ID:           1,
					Provider:     "provider1",
					AccessToken:  "oldAccessToken",
					RefreshToken: "oldRefreshToken",
					TokenType:    "Bearer",
					ExpiresAt:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Hour),
				}
				mockProvider := &provider.Provider{
					ClientID:     "client1",
					ClientSecret: "secret1",
					TokenURL:     "http://token",
				}

				oauthAppRepository.EXPECT().Find(mock.Anything, "owner1", "provider1").Return(mockOAuthApp, nil)
				providerRepository.EXPECT().FindByName(mock.Anything, "provider1").Return(mockProvider, nil)
				tokenSourceFactory.EXPECT().NewTokenSource(mock.Anything, mockProvider, mockOAuthApp).Return(mockTokenSource{
					token: &oauth2.Token{
						AccessToken:  "newAccessToken",
						RefreshToken: "newRefreshToken",
						TokenType:    "Bearer",
						Expiry:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Hour),
					},
					err: nil,
				})
				oauthAppRepository.EXPECT().UpdateByID(mock.Anything, uint(1), mock.Anything).Return(nil)
			},
			expectedError: false,
			expectedResp: &RetrieveAccessTokenResponse{
				AccessToken: "newAccessToken",
			},
		},
		{
			name:       "Failure - Find Error",
			providerID: "provider2",
			ownerID:    "owner2",
			mockSetup: func(oauthAppRepository *MockOAuthAppRepository, providerRepository *MockProviderRepository, tokenSourceFactory *MockTokenSourceFactory) {
				oauthAppRepository.EXPECT().Find(mock.Anything, "owner2", "provider2").Return(nil, errors.New("find error"))
			},
			expectedError: true,
			expectedResp:  nil,
		},
		{
			name:       "Failure - Provider Error",
			providerID: "provider3",
			ownerID:    "owner3",
			mockSetup: func(oauthAppRepository *MockOAuthAppRepository, providerRepository *MockProviderRepository, tokenSourceFactory *MockTokenSourceFactory) {
				oauthAppRepository.EXPECT().Find(mock.Anything, "owner3", "provider3").Return(&oauthapp.OAuthApp{
					ID:          3,
					Provider:    "provider3",
					AccessToken: "oldAccessToken",
					ExpiresAt:   time.Now().Add(-time.Hour),
				}, nil)
				providerRepository.EXPECT().FindByName(mock.Anything, "provider3").Return(nil, errors.New("provider error"))
			},
			expectedError: true,
			expectedResp:  nil,
		},
		{
			name:       "Failure - Token Error",
			providerID: "provider4",
			ownerID:    "owner4",
			mockSetup: func(oauthAppRepository *MockOAuthAppRepository, providerRepository *MockProviderRepository, tokenSourceFactory *MockTokenSourceFactory) {

				mockOAuthApp := &oauthapp.OAuthApp{
					ID:           4,
					Provider:     "provider4",
					AccessToken:  "oldAccessToken",
					RefreshToken: "oldRefreshToken",
					TokenType:    "Bearer",
					ExpiresAt:    time.Now().Add(-time.Hour),
				}
				mockProvider := &provider.Provider{
					ClientID:     "client4",
					ClientSecret: "secret4",
					TokenURL:     "http://token",
				}

				oauthAppRepository.EXPECT().Find(mock.Anything, "owner4", "provider4").Return(mockOAuthApp, nil)
				providerRepository.EXPECT().FindByName(mock.Anything, "provider4").Return(mockProvider, nil)
				tokenSourceFactory.EXPECT().NewTokenSource(mock.Anything, mockProvider, mockOAuthApp).Return(mockTokenSource{
					token: nil,
					err:   errors.New("token error"),
				})
			},
			expectedError: true,
			expectedResp:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oauthAppRepository := NewMockOAuthAppRepository(t)
			providerRepository := NewMockProviderRepository(t)
			tokenSourceFactory := NewMockTokenSourceFactory(t)

			tt.mockSetup(oauthAppRepository, providerRepository, tokenSourceFactory)

			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

			service := NewService(oauthAppRepository, providerRepository, tokenSourceFactory, logger)

			resp, err := service.RetrieveAccessToken(context.Background(), tt.providerID, tt.ownerID)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedResp, resp)
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
