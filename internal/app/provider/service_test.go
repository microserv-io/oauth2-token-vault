package provider

import (
	"context"
	"errors"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/oauth2"
	"testing"
	"time"

	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewService(t *testing.T) {
	mockProviderRepo := NewMockProviderRepository(t)
	mockOAuthAppRepo := NewMockOAuthAppRepository(t)
	mockEncryptor := NewMockEncryptor(t)
	mockOAuth2Client := NewMockOAuth2Client(t)

	service := NewService(mockProviderRepo, mockOAuthAppRepo, mockEncryptor, mockOAuth2Client)
	assert.NotNil(t, service)
}

func TestService_ListProviders(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderRepo *MockProviderRepository)
		expectedError error
		expectedResp  *ListProvidersResponse
	}{
		{
			name: "success",
			mockSetup: func(mockProviderRepo *MockProviderRepository) {
				mockProviderRepo.EXPECT().List(mock.Anything).Return([]*provider.Provider{
					{
						Name:        "provider1",
						AuthURL:     "http://auth.url",
						TokenURL:    "http://token.url",
						Scopes:      []string{"scope1", "scope2"},
						RedirectURL: "http://redirect.url",
						ClientID:    "client_id",
					},
				}, nil)
			},
			expectedError: nil,
			expectedResp: &ListProvidersResponse{
				Providers: []*Provider{
					{
						Name:        "provider1",
						AuthURL:     "http://auth.url",
						TokenURL:    "http://token.url",
						Scopes:      []string{"scope1", "scope2"},
						RedirectURI: "http://redirect.url",
						ClientID:    "client_id",
					},
				},
			},
		},
		{
			name: "error",
			mockSetup: func(mockProviderRepo *MockProviderRepository) {
				mockProviderRepo.EXPECT().List(mock.Anything).Return(nil, errors.New("some error"))
			},
			expectedError: errors.New("failed to list providers: some error"),
			expectedResp:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProviderRepo := NewMockProviderRepository(t)
			service := &Service{
				providerRepository: mockProviderRepo,
			}

			tt.mockSetup(mockProviderRepo)
			resp, err := service.ListProviders(context.Background())
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedResp, resp)
		})
	}
}

func TestService_GetProviderByName(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderRepo *MockProviderRepository)
		expectedError error
		expectedResp  *GetProviderByNameResponse
	}{
		{
			name: "success",
			mockSetup: func(mockProviderRepo *MockProviderRepository) {
				mockProviderRepo.EXPECT().FindByName(mock.Anything, "provider1").Return(&provider.Provider{
					Name:        "provider1",
					AuthURL:     "http://auth.url",
					TokenURL:    "http://token.url",
					Scopes:      []string{"scope1", "scope2"},
					RedirectURL: "http://redirect.url",
					ClientID:    "client_id",
				}, nil)
			},
			expectedError: nil,
			expectedResp: &GetProviderByNameResponse{
				Provider: &Provider{
					Name:        "provider1",
					AuthURL:     "http://auth.url",
					TokenURL:    "http://token.url",
					Scopes:      []string{"scope1", "scope2"},
					RedirectURI: "http://redirect.url",
					ClientID:    "client_id",
				},
			},
		},
		{
			name: "error",
			mockSetup: func(mockProviderRepo *MockProviderRepository) {
				mockProviderRepo.EXPECT().FindByName(mock.Anything, "provider1").Return(nil, errors.New("some error"))
			},
			expectedError: errors.New("some error"),
			expectedResp:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProviderRepo := NewMockProviderRepository(t)
			service := &Service{
				providerRepository: mockProviderRepo,
			}

			tt.mockSetup(mockProviderRepo)

			resp, err := service.GetProviderByName(context.Background(), "provider1")
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedResp, resp)

		})
	}
}

func TestService_CreateProvider(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderRepo *MockProviderRepository)
		input         *CreateInput
		source        string
		expectedError error
		expectedResp  *CreateProviderResponse
	}{
		{
			name: "success",
			mockSetup: func(mockProviderRepo *MockProviderRepository) {
				mockProviderRepo.EXPECT().Create(mock.Anything, mock.Anything).Return(nil)
			},
			input: &CreateInput{
				Name:         "provider1",
				ClientID:     "client_id",
				ClientSecret: "client_secret",
				RedirectURI:  "http://redirect.url",
				AuthURL:      "http://auth.url",
				TokenURL:     "http://token.url",
				Scopes:       []string{"scope1", "scope2"},
			},
			source:        "api",
			expectedError: nil,
			expectedResp: &CreateProviderResponse{
				Provider: &Provider{
					Name:        "provider1",
					AuthURL:     "http://auth.url",
					TokenURL:    "http://token.url",
					Scopes:      []string{"scope1", "scope2"},
					RedirectURI: "http://redirect.url",
					ClientID:    "client_id",
				},
			},
		},
		{
			name: "error",
			mockSetup: func(mockProviderRepo *MockProviderRepository) {
				mockProviderRepo.EXPECT().Create(mock.Anything, mock.Anything).Return(errors.New("some error"))
			},
			input: &CreateInput{
				Name:         "provider1",
				ClientID:     "client_id",
				ClientSecret: "client_secret",
				RedirectURI:  "http://redirect.url",
				AuthURL:      "http://auth.url",
				TokenURL:     "http://token.url",
				Scopes:       []string{"scope1", "scope2"},
			},
			source:        "api",
			expectedError: errors.New("failed to create provider: some error"),
			expectedResp:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProviderRepo := NewMockProviderRepository(t)
			service := &Service{
				providerRepository: mockProviderRepo,
			}

			tt.mockSetup(mockProviderRepo)
			resp, err := service.CreateProvider(context.Background(), tt.input, tt.source)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedResp, resp)
		})
	}
}

func TestService_UpdateProvider(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderRepo *MockProviderRepository, mockEncryptor *MockEncryptor)
		input         *UpdateInput
		expectedError error
		expectedResp  *UpdateProviderResponse
	}{
		{
			name: "success",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockEncryptor *MockEncryptor) {
				mockProviderRepo.EXPECT().FindByName(mock.Anything, "provider1").Return(&provider.Provider{
					Name:        "provider1",
					AuthURL:     "http://auth.url",
					TokenURL:    "http://token.url",
					Scopes:      []string{"scope1", "scope2"},
					RedirectURL: "http://redirect.url",
					ClientID:    "client_id",
					CreatedAt:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				}, nil)
				mockEncryptor.EXPECT().Encrypt("new_client_secret").Return("encrypted_secret", nil)
				mockProviderRepo.EXPECT().Update(mock.Anything, mock.Anything).Return(nil)
			},
			input: &UpdateInput{
				ClientID:     "new_client_id",
				ClientSecret: "new_client_secret",
				RedirectURI:  "http://new.redirect.url",
				AuthURL:      "http://new.auth.url",
				TokenURL:     "http://new.token.url",
				Scopes:       []string{"new_scope1", "new_scope2"},
			},
			expectedError: nil,
			expectedResp: &UpdateProviderResponse{
				Provider: &Provider{
					Name:        "provider1",
					AuthURL:     "http://new.auth.url",
					TokenURL:    "http://new.token.url",
					Scopes:      []string{"new_scope1", "new_scope2"},
					RedirectURI: "http://new.redirect.url",
					ClientID:    "new_client_id",
					CreatedAt:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "error",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockEncryptor *MockEncryptor) {
				mockProviderRepo.EXPECT().FindByName(mock.Anything, "provider1").Return(nil, errors.New("some error"))
			},
			input: &UpdateInput{
				ClientID:     "new_client_id",
				ClientSecret: "new_client_secret",
				RedirectURI:  "http://new.redirect.url",
				AuthURL:      "http://new.auth.url",
				TokenURL:     "http://new.token.url",
				Scopes:       []string{"new_scope1", "new_scope2"},
			},
			expectedError: errors.New("failed to find provider by name: some error"),
			expectedResp:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProviderRepo := NewMockProviderRepository(t)
			mockEncryptor := NewMockEncryptor(t)
			service := &Service{
				providerRepository: mockProviderRepo,
				encryptor:          mockEncryptor,
			}

			tt.mockSetup(mockProviderRepo, mockEncryptor)
			resp, err := service.UpdateProvider(context.Background(), "provider1", tt.input)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedResp, resp)
		})
	}
}

func TestService_DeleteProvider(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderRepo *MockProviderRepository, mockOAuthAppRepo *MockOAuthAppRepository)
		expectedError error
	}{
		{
			name: "success",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockOAuthAppRepo *MockOAuthAppRepository) {
				mockOAuthAppRepo.EXPECT().ListForProvider(mock.Anything, "provider1").Return(nil, nil)
				mockProviderRepo.EXPECT().Delete(mock.Anything, "provider1").Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "error - associated oauth apps",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockOAuthAppRepo *MockOAuthAppRepository) {
				mockOAuthAppRepo.EXPECT().ListForProvider(mock.Anything, "provider1").Return([]*oauthapp.OAuthApp{
					{
						ID:        1,
						Provider:  "provider1",
						Scopes:    []string{"scope1", "scope2"},
						CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				}, nil)
			},
			expectedError: errors.New("provider has associated oauth apps"),
		},
		{
			name: "error - delete failure",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockOAuthAppRepo *MockOAuthAppRepository) {
				mockOAuthAppRepo.EXPECT().ListForProvider(mock.Anything, "provider1").Return(nil, nil)
				mockProviderRepo.EXPECT().Delete(mock.Anything, "provider1").Return(errors.New("some error"))
			},
			expectedError: errors.New("failed to delete provider: some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProviderRepo := NewMockProviderRepository(t)
			mockOAuthAppRepo := NewMockOAuthAppRepository(t)
			service := &Service{
				providerRepository: mockProviderRepo,
				oauthAppRepository: mockOAuthAppRepo,
			}

			tt.mockSetup(mockProviderRepo, mockOAuthAppRepo)
			err := service.DeleteProvider(context.Background(), "provider1")
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_ExchangeAuthorizationCode(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderRepo *MockProviderRepository, mockOAuth2Client *MockOAuth2Client, oauthAppRepo *MockOAuthAppRepository)
		input         *ExchangeAuthorizationCodeInput
		expectedError error
	}{
		{
			name: "success",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockOAuth2Client *MockOAuth2Client, oauthAppRepo *MockOAuthAppRepository) {
				mockProvider := &provider.Provider{
					Name:         "provider1",
					AuthURL:      "http://auth.url",
					TokenURL:     "http://token.url",
					Scopes:       []string{"scope1", "scope2"},
					RedirectURL:  "http://redirect.url",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
				}
				mockProviderRepo.EXPECT().FindByName(mock.Anything, "provider1").Return(mockProvider, nil)

				mockOAuth2Client.EXPECT().Exchange(mock.Anything, &oauth2.Config{
					ClientID:     mockProvider.ClientID,
					ClientSecret: mockProvider.ClientSecret,
					AuthURL:      mockProvider.AuthURL,
					TokenURL:     mockProvider.TokenURL,
					RedirectURL:  mockProvider.RedirectURL,
					Scopes:       mockProvider.Scopes,
				}, "code").Return(&oauth2.Token{
					TokenType:    "token_type",
					AccessToken:  "access_token",
					RefreshToken: "refresh_token",
					ExpiresAt:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				}, nil)

				oauthAppRepo.EXPECT().Create(mock.Anything, &oauthapp.OAuthApp{
					Provider:     mockProvider.Name,
					OwnerID:      "owner1",
					Scopes:       mockProvider.Scopes,
					AccessToken:  "access_token",
					RefreshToken: "refresh_token",
					TokenType:    "token_type",
					ExpiresAt:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				}).Return(nil)
			},
			input: &ExchangeAuthorizationCodeInput{
				Provider: "provider1",
				OwnerID:  "owner1",
				Code:     "code",
			},
			expectedError: nil,
		},
		{
			name: "error - find provider failure",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockOAuth2Client *MockOAuth2Client, oauthAppRepo *MockOAuthAppRepository) {
				mockProviderRepo.EXPECT().FindByName(mock.Anything, "provider1").Return(nil, errors.New("some error"))
			},
			input: &ExchangeAuthorizationCodeInput{
				Provider: "provider1",
				OwnerID:  "owner1",
				Code:     "code",
			},
			expectedError: errors.New("failed to find providerObj by name: some error"),
		},
		{
			name: "error - exchange failure",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockOAuth2Client *MockOAuth2Client, mockOAuthAppRepo *MockOAuthAppRepository) {
				mockProvider := &provider.Provider{
					Name:         "provider1",
					AuthURL:      "http://auth.url",
					TokenURL:     "http://token.url",
					Scopes:       []string{"scope1", "scope2"},
					RedirectURL:  "http://redirect.url",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
				}
				mockProviderRepo.EXPECT().FindByName(mock.Anything, "provider1").Return(mockProvider, nil)
				mockOAuth2Client.EXPECT().Exchange(mock.Anything, &oauth2.Config{
					ClientID:     mockProvider.ClientID,
					ClientSecret: mockProvider.ClientSecret,
					AuthURL:      mockProvider.AuthURL,
					TokenURL:     mockProvider.TokenURL,
					RedirectURL:  mockProvider.RedirectURL,
					Scopes:       mockProvider.Scopes,
				}, "code").Return(nil, errors.New("some error"))
			},
			input: &ExchangeAuthorizationCodeInput{
				Provider: "provider1",
				OwnerID:  "owner1",
				Code:     "code",
			},
			expectedError: errors.New("failed to exchange authorization code: some error"),
		},
		{
			name: "error - create oauth app failure",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockOAuth2Client *MockOAuth2Client, mockOAuthAppRepo *MockOAuthAppRepository) {
				mockProvider := &provider.Provider{
					Name:         "provider1",
					AuthURL:      "http://auth.url",
					TokenURL:     "http://token.url",
					Scopes:       []string{"scope1", "scope2"},
					RedirectURL:  "http://redirect.url",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
				}
				mockProviderRepo.EXPECT().FindByName(mock.Anything, "provider1").Return(mockProvider, nil)

				mockOAuth2Client.EXPECT().Exchange(mock.Anything, &oauth2.Config{
					ClientID:     mockProvider.ClientID,
					ClientSecret: mockProvider.ClientSecret,
					AuthURL:      mockProvider.AuthURL,
					TokenURL:     mockProvider.TokenURL,
					RedirectURL:  mockProvider.RedirectURL,
					Scopes:       mockProvider.Scopes,
				}, "code").Return(&oauth2.Token{
					TokenType:    "token_type",
					AccessToken:  "access_token",
					RefreshToken: "refresh_token",
					ExpiresAt:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				}, nil)

				mockOAuthAppRepo.EXPECT().Create(mock.Anything, &oauthapp.OAuthApp{
					Provider:     mockProvider.Name,
					OwnerID:      "owner1",
					Scopes:       mockProvider.Scopes,
					AccessToken:  "access_token",
					RefreshToken: "refresh_token",
					TokenType:    "token_type",
					ExpiresAt:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				}).Return(errors.New("some error"))
			},
			input: &ExchangeAuthorizationCodeInput{
				Provider: "provider1",
				OwnerID:  "owner1",
				Code:     "code",
			},
			expectedError: errors.New("failed to create oauth app: some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProviderRepo := NewMockProviderRepository(t)
			mockOAuth2Client := NewMockOAuth2Client(t)
			mockOAuthAppRepo := NewMockOAuthAppRepository(t)
			service := &Service{
				providerRepository: mockProviderRepo,
				oauth2Client:       mockOAuth2Client,
				oauthAppRepository: mockOAuthAppRepo,
			}

			tt.mockSetup(mockProviderRepo, mockOAuth2Client, mockOAuthAppRepo)
			err := service.ExchangeAuthorizationCode(context.Background(), tt.input)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
