package v1

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/microserv-io/oauth2-token-vault/internal/app/provider"
	"github.com/microserv-io/oauth2-token-vault/pkg/proto/oauthcredentials/v1"
	"github.com/stretchr/testify/assert"
)

func TestProviderServiceGRPC_ListProviders(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderService *MockProviderService)
		expectedError error
		expectedResp  *oauthcredentials.ListProvidersResponse
	}{
		{
			name: "success",
			mockSetup: func(mockProviderService *MockProviderService) {
				mockProviderService.EXPECT().ListProviders(mock.Anything).Return(&provider.ListProvidersResponse{
					Providers: []*provider.Provider{
						{Name: "provider1", AuthURL: "http://auth.url", TokenURL: "http://token.url", Scopes: []string{"scope1"}, ClientID: "client1"},
					},
				}, nil)
			},
			expectedError: nil,
			expectedResp: &oauthcredentials.ListProvidersResponse{
				OauthProviders: []*oauthcredentials.OAuthProvider{
					{Name: "provider1", AuthUrl: "http://auth.url", TokenUrl: "http://token.url", Scopes: []string{"scope1"}, ClientId: "client1"},
				},
			},
		},
		{
			name: "error",
			mockSetup: func(mockProviderService *MockProviderService) {
				mockProviderService.EXPECT().ListProviders(mock.Anything).Return(nil, errors.New("some error"))
			},
			expectedError: errors.New("failed to list providers: some error"),
			expectedResp:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProviderService := NewMockProviderService(t)
			tt.mockSetup(mockProviderService)

			s := NewProviderServiceGRPC(mockProviderService)

			resp, err := s.ListProviders(context.Background(), &oauthcredentials.ListProvidersRequest{})

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedResp, resp)
		})
	}
}

func TestProviderServiceGRPC_CreateProvider(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderService *MockProviderService)
		expectedError error
		expectedResp  *oauthcredentials.CreateProviderResponse
	}{
		{
			name: "success",
			mockSetup: func(mockProviderService *MockProviderService) {
				mockProviderService.EXPECT().CreateProvider(mock.Anything, &provider.CreateInput{
					Name:    "provider1",
					AuthURL: "http://auth.url",
				}, "api").Return(&provider.CreateProviderResponse{
					Provider: &provider.Provider{
						Name: "provider1", AuthURL: "http://auth.url", TokenURL: "http://token.url", Scopes: []string{"scope1"}, ClientID: "client1",
					},
				}, nil)
			},
			expectedError: nil,
			expectedResp: &oauthcredentials.CreateProviderResponse{
				OauthProvider: &oauthcredentials.OAuthProvider{
					Name: "provider1", AuthUrl: "http://auth.url", TokenUrl: "http://token.url", Scopes: []string{"scope1"}, ClientId: "client1",
				},
			},
		},
		{
			name: "error",
			mockSetup: func(mockProviderService *MockProviderService) {
				mockProviderService.EXPECT().CreateProvider(mock.Anything, &provider.CreateInput{
					Name:    "provider1",
					AuthURL: "http://auth.url",
				}, "api").Return(nil, errors.New("some error"))
			},
			expectedError: errors.New("failed to create provider: some error"),
			expectedResp:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProviderService := NewMockProviderService(t)
			tt.mockSetup(mockProviderService)

			s := NewProviderServiceGRPC(mockProviderService)

			resp, err := s.CreateProvider(context.TODO(), &oauthcredentials.CreateProviderRequest{
				Name:    "provider1",
				AuthUrl: "http://auth.url",
			})

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}

func TestProviderServiceGRPC_UpdateProvider(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderService *MockProviderService)
		expectedError error
		expectedResp  *oauthcredentials.UpdateProviderResponse
	}{
		{
			name: "success",
			mockSetup: func(mockProviderService *MockProviderService) {
				mockProviderService.EXPECT().UpdateProvider(mock.Anything, "provider1", &provider.UpdateInput{
					AuthURL: "http://auth.url",
				}).Return(&provider.UpdateProviderResponse{
					Provider: &provider.Provider{
						Name: "provider1", AuthURL: "http://auth.url", TokenURL: "http://token.url", Scopes: []string{"scope1"}, ClientID: "client1",
					},
				}, nil)
			},
			expectedError: nil,
			expectedResp: &oauthcredentials.UpdateProviderResponse{
				OauthProvider: &oauthcredentials.OAuthProvider{
					Name: "provider1", AuthUrl: "http://auth.url", TokenUrl: "http://token.url", Scopes: []string{"scope1"}, ClientId: "client1",
				},
			},
		},
		{
			name: "error",
			mockSetup: func(mockProviderService *MockProviderService) {
				mockProviderService.EXPECT().UpdateProvider(mock.Anything, "provider1", &provider.UpdateInput{
					AuthURL: "http://auth.url",
				}).Return(nil, errors.New("some error"))
			},
			expectedError: errors.New("failed to update provider: some error"),
			expectedResp:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProviderService := NewMockProviderService(t)
			tt.mockSetup(mockProviderService)

			s := NewProviderServiceGRPC(mockProviderService)

			resp, err := s.UpdateProvider(context.TODO(), &oauthcredentials.UpdateProviderRequest{
				Name:    "provider1",
				AuthUrl: "http://auth.url",
			})

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}

func TestProviderServiceGRPC_DeleteProvider(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderService *MockProviderService)
		expectedError error
	}{
		{
			name: "success",
			mockSetup: func(mockProviderService *MockProviderService) {
				mockProviderService.EXPECT().DeleteProvider(mock.Anything, "provider1").Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "error",
			mockSetup: func(mockProviderService *MockProviderService) {
				mockProviderService.EXPECT().DeleteProvider(mock.Anything, "provider1").Return(errors.New("some error"))
			},
			expectedError: errors.New("failed to delete provider: some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProviderService := NewMockProviderService(t)
			tt.mockSetup(mockProviderService)

			s := NewProviderServiceGRPC(mockProviderService)

			_, err := s.DeleteProvider(context.TODO(), &oauthcredentials.DeleteProviderRequest{
				Name: "provider1",
			})

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProviderServiceGRPC_ExchangeAuthorizationCode(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockProviderService *MockProviderService)
		expectedError error
	}{
		{
			name: "success",
			mockSetup: func(mockProviderService *MockProviderService) {
				mockProviderService.EXPECT().ExchangeAuthorizationCode(mock.Anything, &provider.ExchangeAuthorizationCodeInput{
					Provider: "provider1",
				}).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "error",
			mockSetup: func(mockProviderService *MockProviderService) {
				mockProviderService.EXPECT().ExchangeAuthorizationCode(mock.Anything, &provider.ExchangeAuthorizationCodeInput{
					Provider: "provider1",
				}).Return(errors.New("some error"))
			},
			expectedError: errors.New("failed to exchange authorization code: some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProviderService := NewMockProviderService(t)
			tt.mockSetup(mockProviderService)

			s := NewProviderServiceGRPC(mockProviderService)

			_, err := s.ExchangeAuthorizationCode(context.TODO(), &oauthcredentials.ExchangeAuthorizationCodeRequest{
				Provider: "provider1",
			})

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
