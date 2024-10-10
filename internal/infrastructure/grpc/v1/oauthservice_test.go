package v1

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/microserv-io/oauth-credentials-server/internal/app/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
	"github.com/stretchr/testify/assert"
)

func TestService_ListOAuths(t *testing.T) {

	tests := []struct {
		name          string
		mockSetup     func(oauthAppService *MockOAuthAppService)
		request       *oauthcredentials.ListOAuthsRequest
		expectedError error
		expectedResp  *oauthcredentials.ListOAuthsResponse
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
			request: &oauthcredentials.ListOAuthsRequest{Owner: "owner1"},
			expectedResp: &oauthcredentials.ListOAuthsResponse{
				Oauths: []*oauthcredentials.OAuth{
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
			request:       &oauthcredentials.ListOAuthsRequest{Owner: "owner1"},
			expectedResp:  nil,
			expectedError: errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOAuthAppService := NewMockOAuthAppService(t)

			tt.mockSetup(mockOAuthAppService)

			service := NewOAuthAppServiceGRPC(mockOAuthAppService)

			stream := NewMockListOAuthsStream(t)

			stream.EXPECT().Context().Return(context.Background())

			if tt.expectedResp != nil {
				stream.EXPECT().Send(tt.expectedResp).Return(nil)
			}

			err := service.ListOAuths(tt.request, stream)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
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
				Oauth: &oauthcredentials.OAuth{
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
			expectedError: errors.New("some error"),
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
