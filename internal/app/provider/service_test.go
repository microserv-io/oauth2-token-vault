package provider

import (
	"context"
	"errors"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/oauth2"
	"net/url"
	"testing"
	"time"

	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/provider"
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
		input         *CreateProviderRequest
		source        string
		expectedError error
		expectedResp  *CreateProviderResponse
	}{
		{
			name: "success",
			mockSetup: func(mockProviderRepo *MockProviderRepository) {
				mockProviderRepo.EXPECT().Create(mock.Anything, mock.Anything).Return(nil)
			},
			input: &CreateProviderRequest{
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
			input: &CreateProviderRequest{
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
		input         *UpdateProviderRequest
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
			input: &UpdateProviderRequest{
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
			input: &UpdateProviderRequest{
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
			err := service.DeleteProvider(context.Background(), &DeleteProviderRequest{
				Name:                     "provider1",
				DeleteConnectedOAuthApps: false,
			})
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_GetAuthorizationURL(t *testing.T) {
	tests := []struct {
		name          string
		input         *GetAuthorizationURLRequest
		state         string
		mockSetup     func(providerRepository *MockProviderRepository, oauthClient *MockOAuth2Client)
		expectedError bool
		expectedResp  *GetAuthorizationURLResponse
	}{
		{
			name:  "Success",
			input: &GetAuthorizationURLRequest{Provider: "provider1", State: "state1"},
			state: "state1",
			mockSetup: func(providerRepository *MockProviderRepository, oauthClient *MockOAuth2Client) {
				providerRepository.EXPECT().FindByName(mock.Anything, "provider1").Return(&provider.Provider{
					ClientID:     "client1",
					ClientSecret: "secret1",
					AuthURL:      "http://auth",
					RedirectURL:  "http://localhost",
					Scopes:       []string{"scope1"},
				}, nil)
				oauthClient.EXPECT().GetAuthorizationURL(&oauth2.Config{
					ClientID:     "client1",
					ClientSecret: "secret1",
					AuthURL:      "http://auth",
					RedirectURL:  "http://localhost",
					Scopes:       []string{"scope1"},
				}, "state1").Return("http://auth?client_id=client1&redirect_uri=http%3A%2F%2Flocalhost&response_type=code&scope=scope1&state=state1", nil)
			},
			expectedError: false,
			expectedResp: &GetAuthorizationURLResponse{
				URL: &url.URL{
					Scheme:   "http",
					Host:     "auth",
					RawQuery: "client_id=client1&redirect_uri=http%3A%2F%2Flocalhost&response_type=code&scope=scope1&state=state1",
				},
			},
		},
		{
			name:  "Failure",
			input: &GetAuthorizationURLRequest{Provider: "provider2", State: "state2"},
			state: "state2",
			mockSetup: func(providerRepository *MockProviderRepository, oauthClient *MockOAuth2Client) {
				providerRepository.EXPECT().FindByName(mock.Anything, "provider2").Return(nil, errors.New("database error"))
			},
			expectedError: true,
			expectedResp:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			providerRepository := NewMockProviderRepository(t)
			oauthAppRepository := NewMockOAuthAppRepository(t)
			mockEncryptor := NewMockEncryptor(t)
			mockOAuth2Client := NewMockOAuth2Client(t)

			tt.mockSetup(providerRepository, mockOAuth2Client)

			service := NewService(providerRepository, oauthAppRepository, mockEncryptor, mockOAuth2Client)

			resp, err := service.GetAuthorizationURL(context.Background(), tt.input)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedResp, resp)
		})
	}
}

func TestService_ExchangeAuthorizationCode(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderRepo *MockProviderRepository, mockOAuth2Client *MockOAuth2Client, oauthAppRepo *MockOAuthAppRepository, mockEncryptor *MockEncryptor)
		input         *ExchangeAuthorizationCodeRequest
		expectedError error
	}{
		{
			name: "success",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockOAuth2Client *MockOAuth2Client, oauthAppRepo *MockOAuthAppRepository, mockEncryptor *MockEncryptor) {
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

				mockEncryptor.EXPECT().Encrypt("access_token").Return("encrypted_access_token", nil)
				mockEncryptor.EXPECT().Encrypt("refresh_token").Return("encrypted_refresh_token", nil)

				oauthAppRepo.EXPECT().Create(mock.Anything, &oauthapp.OAuthApp{
					Provider:     mockProvider.Name,
					OwnerID:      "owner1",
					Scopes:       mockProvider.Scopes,
					AccessToken:  "encrypted_access_token",
					RefreshToken: "encrypted_refresh_token",
					TokenType:    "token_type",
					ExpiresAt:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				}).Return(nil)

			},
			input: &ExchangeAuthorizationCodeRequest{
				Provider: "provider1",
				OwnerID:  "owner1",
				Code:     "code",
			},
			expectedError: nil,
		},
		{
			name: "error - find provider failure",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockOAuth2Client *MockOAuth2Client, oauthAppRepo *MockOAuthAppRepository, mockEncryptor *MockEncryptor) {
				mockProviderRepo.EXPECT().FindByName(mock.Anything, "provider1").Return(nil, errors.New("some error"))
			},
			input: &ExchangeAuthorizationCodeRequest{
				Provider: "provider1",
				OwnerID:  "owner1",
				Code:     "code",
			},
			expectedError: errors.New("failed to find providerObj by name: some error"),
		},
		{
			name: "error - exchange failure",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockOAuth2Client *MockOAuth2Client, mockOAuthAppRepo *MockOAuthAppRepository, mockEncryptor *MockEncryptor) {
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
			input: &ExchangeAuthorizationCodeRequest{
				Provider: "provider1",
				OwnerID:  "owner1",
				Code:     "code",
			},
			expectedError: errors.New("failed to exchange authorization code: some error"),
		},
		{
			name: "error - create oauth app failure",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockOAuth2Client *MockOAuth2Client, mockOAuthAppRepo *MockOAuthAppRepository, mockEncryptor *MockEncryptor) {
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

				mockEncryptor.EXPECT().Encrypt("access_token").Return("encrypted_access_token", nil)
				mockEncryptor.EXPECT().Encrypt("refresh_token").Return("encrypted_refresh_token", nil)

				mockOAuthAppRepo.EXPECT().Create(mock.Anything, &oauthapp.OAuthApp{
					Provider:     mockProvider.Name,
					OwnerID:      "owner1",
					Scopes:       mockProvider.Scopes,
					AccessToken:  "encrypted_access_token",
					RefreshToken: "encrypted_refresh_token",
					TokenType:    "token_type",
					ExpiresAt:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				}).Return(errors.New("some error"))
			},
			input: &ExchangeAuthorizationCodeRequest{
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
			mockEncryptor := NewMockEncryptor(t)
			service := &Service{
				providerRepository: mockProviderRepo,
				oauth2Client:       mockOAuth2Client,
				oauthAppRepository: mockOAuthAppRepo,
				encryptor:          mockEncryptor,
			}

			tt.mockSetup(mockProviderRepo, mockOAuth2Client, mockOAuthAppRepo, mockEncryptor)
			err := service.ExchangeAuthorizationCode(context.Background(), tt.input)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_SyncProviders(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderRepo *MockProviderRepository, mockOAuthAppRepo *MockOAuthAppRepository, encryptor *MockEncryptor)
		input         *SyncProviderRequest
		expectedError error
	}{
		{
			name: "success - update",
			mockSetup: func(mockProviderRepo *MockProviderRepository, mockOAuthAppRepo *MockOAuthAppRepository, mockEncryptor *MockEncryptor) {
				mockProviderRepo.EXPECT().List(mock.Anything).Return([]*provider.Provider{
					{
						Name:   "provider1",
						Source: "config",
					},
					{
						Name:   "provider2",
						Source: "config",
					},
					{
						Name:   "provider3",
						Source: "config",
					},
					{
						Name:   "provider4",
						Source: "api",
					},
				}, nil)

				mockProviderRepo.EXPECT().FindByName(mock.Anything, "provider1").Return(&provider.Provider{
					Name:     "provider1",
					ClientID: "old_client_id",
				}, nil)

				mockEncryptor.EXPECT().Encrypt("client_secret1").Return("encrypted_secret", nil)
				mockEncryptor.EXPECT().Encrypt("client_secret2").Return("encrypted_secret", nil)

				mockProviderRepo.EXPECT().Update(mock.Anything, &provider.Provider{
					Name:         "provider1",
					ClientID:     "client_id",
					ClientSecret: "encrypted_secret",
					RedirectURL:  "http://redirect.url",
					AuthURL:      "http://auth.url",
					TokenURL:     "http://token.url",
					Scopes:       []string{"scope1", "scope2"},
					Source:       "config",
				}).Return(nil)

				mockProviderRepo.EXPECT().FindByName(mock.Anything, "provider2").Return(nil, errors.New("some error"))
				mockProviderRepo.EXPECT().Create(mock.Anything, &provider.Provider{
					Name:         "provider2",
					ClientID:     "client_id",
					ClientSecret: "encrypted_secret",
					RedirectURL:  "http://redirect.url",
					AuthURL:      "http://auth.url",
					TokenURL:     "http://token.url",
					Scopes:       []string{"scope1", "scope2"},
					Source:       "config",
				}).Return(nil)

				mockProviderRepo.EXPECT().Delete(mock.Anything, "provider3").Return(nil)
			},
			input: &SyncProviderRequest{
				Providers: []*SyncProvider{
					{
						Name:         "provider1",
						ClientID:     "client_id",
						ClientSecret: "client_secret1",
						RedirectURI:  "http://redirect.url",
						AuthURL:      "http://auth.url",
						TokenURL:     "http://token.url",
						Scopes:       []string{"scope1", "scope2"},
					},
					{
						Name:         "provider2",
						ClientID:     "client_id",
						ClientSecret: "client_secret2",
						RedirectURI:  "http://redirect.url",
						AuthURL:      "http://auth.url",
						TokenURL:     "http://token.url",
						Scopes:       []string{"scope1", "scope2"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProviderRepo := NewMockProviderRepository(t)
			mockOAuthRepo := NewMockOAuthAppRepository(t)
			mockEncryptor := NewMockEncryptor(t)

			service := &Service{
				providerRepository: mockProviderRepo,
				oauthAppRepository: mockOAuthRepo,
				encryptor:          mockEncryptor,
			}

			tt.mockSetup(mockProviderRepo, mockOAuthRepo, mockEncryptor)
			err := service.SyncProviders(context.Background(), tt.input)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
