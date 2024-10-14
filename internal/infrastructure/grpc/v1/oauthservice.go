package v1

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/app/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
)

var _ oauthcredentials.OAuthServiceServer = &OAuthAppServiceGRPC{}

type OAuthAppService interface {
	ListOAuthAppsForOwner(ctx context.Context, ownerID string) (*oauthapp.ListOAuthAppsForOwnerResponse, error)
	GetOAuthForProviderAndOwner(ctx context.Context, providerID, ownerID string) (*oauthapp.GetOAuthForProviderAndOwnerResponse, error)
	RetrieveAccessToken(ctx context.Context, providerID, ownerID string) (*oauthapp.RetrieveAccessTokenResponse, error)
}

type OAuthAppServiceGRPC struct {
	oauthcredentials.UnimplementedOAuthServiceServer

	oauthAppService OAuthAppService
}

func NewOAuthAppServiceGRPC(
	oauthAppService OAuthAppService,
) *OAuthAppServiceGRPC {

	service := OAuthAppServiceGRPC{
		oauthAppService: oauthAppService,
	}

	return &service
}

func (s OAuthAppServiceGRPC) ListOAuths(ctx context.Context, request *oauthcredentials.ListOAuthsForOwnerRequest) (*oauthcredentials.ListOAuthsForOwnerResponse, error) {
	oauthApps, err := s.oauthAppService.ListOAuthAppsForOwner(ctx, request.GetOwner())
	if err != nil {
		return nil, fmt.Errorf("could not list oauth apps: %w", err)
	}

	var oauthAppsResponse []*oauthcredentials.OAuthApp
	for _, oauthApp := range oauthApps.Apps {
		oauthAppsResponse = append(oauthAppsResponse, &oauthcredentials.OAuthApp{
			Id:       oauthApp.ID,
			Owner:    oauthApp.OwnerID,
			Provider: oauthApp.ProviderID,
			Scopes:   oauthApp.Scopes,
		})
	}

	return &oauthcredentials.ListOAuthsForOwnerResponse{
		OauthApps: oauthAppsResponse,
	}, nil
}

func (s OAuthAppServiceGRPC) GetOAuthByProvider(ctx context.Context, request *oauthcredentials.GetOAuthByProviderRequest) (*oauthcredentials.GetOAuthByProviderResponse, error) {
	resp, err := s.oauthAppService.GetOAuthForProviderAndOwner(ctx, request.GetProvider(), request.GetOwner())

	if err != nil {
		return nil, fmt.Errorf("could not get oauth app: %w", err)
	}

	return &oauthcredentials.GetOAuthByProviderResponse{
		OauthApp: &oauthcredentials.OAuthApp{
			Id:       resp.App.ID,
			Owner:    resp.App.OwnerID,
			Provider: resp.App.ProviderID,
			Scopes:   resp.App.Scopes,
		},
	}, nil
}

func (s OAuthAppServiceGRPC) GetOAuthCredentialByProvider(ctx context.Context, request *oauthcredentials.GetOAuthCredentialByProviderRequest) (*oauthcredentials.GetOAuthCredentialByProviderResponse, error) {
	resp, err := s.oauthAppService.RetrieveAccessToken(ctx, request.GetProvider(), request.GetOwner())
	if err != nil {
		return nil, fmt.Errorf("could not get oauth credentials: %w", err)
	}

	return &oauthcredentials.GetOAuthCredentialByProviderResponse{
		AccessToken: resp.AccessToken,
	}, nil
}
