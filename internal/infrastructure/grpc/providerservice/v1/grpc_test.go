package providerservice

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/microserv-io/oauth-credentials-server/internal/app/provider"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
	"github.com/stretchr/testify/assert"
)

func TestService_ListProviders(t *testing.T) {
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

			s := NewService(mockProviderService)
			stream := NewMockListProviderStream(t)

			stream.EXPECT().Context().Return(nil)

			if tt.expectedResp != nil {
				stream.EXPECT().Send(tt.expectedResp).Return(nil)
			}

			err := s.ListProviders(&oauthcredentials.ListProvidersRequest{}, stream)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}