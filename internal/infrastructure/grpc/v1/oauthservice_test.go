package v1

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/microserv-io/oauth2-token-vault/internal/app/oauthapp"
	"github.com/microserv-io/oauth2-token-vault/pkg/proto/oauthcredentials/v1"
	"github.com/stretchr/testify/assert"
)

func TestService_ListOAuthsForOwner(t *testing.T) {

	tests := []struct {
		name          string
		mockSetup     func(oauthAppService *MockOAuthAppService)
		request       *oauthcredentials.ListOAuthsForOwnerRequest
		expectedError error
		expectedResp  *oauthcredentials.ListOAuthsForOwnerResponse
	}{
		{
			name: "success",
			mockSetup: func(mockOAuthAppService *MockOAuthAppService) {
				mockOAuthAppService.EXPECT().
					ListOAuthAppsForOwner(mock.Anything, "owner1").
					Return(&oauthapp.ListOAuthAppsForOwnerResponse{
						Apps: []*oauthapp.OAuthApp{
							{ID: "1", OwnerID: "owner1", ProviderID: "provider1", Scopes: []string{"scope1"}},
						},
					}, nil)
			},
			request: &oauthcredentials.ListOAuthsForOwnerRequest{Owner: "owner1"},
			expectedResp: &oauthcredentials.ListOAuthsForOwnerResponse{
				OauthApps: []*oauthcredentials.OAuthApp{
					{Id: "1", Owner: "owner1", Provider: "provider1", Scopes: []string{"scope1"}},
				},
			},
			expectedError: nil,
		},
		{
			name: "error",
			mockSetup: func(mockOAuthAppService *MockOAuthAppService) {
				mockOAuthAppService.EXPECT().
					ListOAuthAppsForOwner(mock.Anything, "owner1").
					Return(nil, errors.New("some error"))
			},
			request:       &oauthcredentials.ListOAuthsForOwnerRequest{Owner: "owner1"},
			expectedResp:  nil,
			expectedError: errors.New("could not list oauth apps: some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOAuthAppService := NewMockOAuthAppService(t)

			tt.mockSetup(mockOAuthAppService)

			service := NewOAuthAppServiceGRPC(mockOAuthAppService)

			resp, err := service.ListOAuths(context.Background(), tt.request)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedResp, resp)
		})
	}
}

func TestService_GetOAuthByProvider(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(oauthAppService *MockOAuthAppService)
		request       *oauthcredentials.GetOAuthByProviderRequest
		expectedError error
		expectedResp  *oauthcredentials.GetOAuthByProviderResponse
	}{
		{
			name: "success",
			mockSetup: func(mockOAuthAppService *MockOAuthAppService) {
				mockOAuthAppService.EXPECT().
					GetOAuthForProviderAndOwner(mock.Anything, "provider1", "owner1").
					Return(&oauthapp.GetOAuthForProviderAndOwnerResponse{
						App: &oauthapp.OAuthApp{ID: "1", OwnerID: "owner1", ProviderID: "provider1", Scopes: []string{"scope1"}},
					}, nil)
			},
			request: &oauthcredentials.GetOAuthByProviderRequest{Provider: "provider1", Owner: "owner1"},
			expectedResp: &oauthcredentials.GetOAuthByProviderResponse{
				OauthApp: &oauthcredentials.OAuthApp{
					Id:       "1",
					Owner:    "owner1",
					Provider: "provider1",
					Scopes:   []string{"scope1"},
				},
			},
			expectedError: nil,
		},
		{
			name: "error",
			mockSetup: func(mockOAuthAppService *MockOAuthAppService) {
				mockOAuthAppService.EXPECT().
					GetOAuthForProviderAndOwner(mock.Anything, "provider1", "owner1").
					Return(nil, errors.New("some error"))
			},
			request:       &oauthcredentials.GetOAuthByProviderRequest{Provider: "provider1", Owner: "owner1"},
			expectedResp:  nil,
			expectedError: errors.New("could not get oauth app: some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockOAuthAppService := NewMockOAuthAppService(t)

			tt.mockSetup(mockOAuthAppService)

			service := NewOAuthAppServiceGRPC(mockOAuthAppService)

			resp, err := service.GetOAuthByProvider(context.Background(), tt.request)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}

func TestService_GetOAuthCredentialByProvider(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(mockOAuthAppService *MockOAuthAppService)
		request       *oauthcredentials.GetOAuthCredentialByProviderRequest
		expectedError error
		expectedResp  *oauthcredentials.GetOAuthCredentialByProviderResponse
	}{
		{
			name: "success",
			mockSetup: func(mockOAuthAppService *MockOAuthAppService) {
				mockOAuthAppService.EXPECT().
					RetrieveAccessToken(mock.Anything, "provider1", "owner1").
					Return(&oauthapp.RetrieveAccessTokenResponse{AccessToken: "token1"}, nil)
			},
			request: &oauthcredentials.GetOAuthCredentialByProviderRequest{Provider: "provider1", Owner: "owner1"},
			expectedResp: &oauthcredentials.GetOAuthCredentialByProviderResponse{
				AccessToken: "token1",
			},
			expectedError: nil,
		},
		{
			name: "error",
			mockSetup: func(mockOAuthAppService *MockOAuthAppService) {
				mockOAuthAppService.EXPECT().
					RetrieveAccessToken(mock.Anything, "provider1", "owner1").
					Return(nil, errors.New("some error"))
			},
			request:       &oauthcredentials.GetOAuthCredentialByProviderRequest{Provider: "provider1", Owner: "owner1"},
			expectedResp:  nil,
			expectedError: errors.New("could not get oauth credentials: some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOAuthAppService := NewMockOAuthAppService(t)
			tt.mockSetup(mockOAuthAppService)

			service := NewOAuthAppServiceGRPC(mockOAuthAppService)
			resp, err := service.GetOAuthCredentialByProvider(context.Background(), tt.request)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}
