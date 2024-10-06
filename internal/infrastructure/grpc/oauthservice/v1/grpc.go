package v1

import (
	"context"
	"github.com/microserv-io/oauth-credentials-server/internal/app/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
)

var _ oauthcredentials.OAuthServiceServer = &Service{}

type Service struct {
	oauthcredentials.UnimplementedOAuthServiceServer

	oauthAppService *oauthapp.Service
}

func NewService(
	oauthAppService *oauthapp.Service,
) *Service {

	service := Service{
		oauthAppService: oauthAppService,
	}

	return &service
}

func (s Service) ListOAuths(request *oauthcredentials.ListOAuthsRequest, server oauthcredentials.OAuthService_ListOAuthsServer) error {
	oauthApps, err := s.oauthAppService.ListOAuthsForOwner(server.Context(), request.GetOwner())
	if err != nil {
		return err
	}

	for _, oauthApp := range oauthApps {

		if err := server.Send(&oauthcredentials.ListOAuthsResponse{Oauths: []*oauthcredentials.OAuth{
			{
				Id:       oauthApp.ID,
				Owner:    oauthApp.OwnerID,
				Provider: oauthApp.ProviderID,
				Scopes:   oauthApp.Scopes,
			},
		}}); err != nil {
			return err
		}
	}

	return nil
}

func (s Service) GetOAuthByID(ctx context.Context, request *oauthcredentials.GetOAuthByIDRequest) (*oauthcredentials.GetOAuthByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetOAuthByProvider(ctx context.Context, request *oauthcredentials.GetOAuthByProviderRequest) (*oauthcredentials.GetOAuthByProviderResponse, error) {
	oauthApp, err := s.oauthAppService.GetOAuthForProviderAndOwner(ctx, request.GetProvider(), request.GetOwner())

	if err != nil {
		return nil, err
	}

	return &oauthcredentials.GetOAuthByProviderResponse{
		Oauth: &oauthcredentials.OAuth{
			Id:       oauthApp.ID,
			Owner:    oauthApp.OwnerID,
			Provider: oauthApp.ProviderID,
			Scopes:   oauthApp.Scopes,
		},
	}, nil
}

func (s Service) GetOAuthCredentialByProvider(ctx context.Context, request *oauthcredentials.GetOAuthCredentialByProviderRequest) (*oauthcredentials.GetOAuthCredentialByProviderResponse, error) {
	//TODO implement me
	panic("implement	me")
}
