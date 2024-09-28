package v1

import (
	"context"
	"github.com/microserv-io/oauth-credentials-server/internal/usecase"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
)

type Service struct {
	oauthcredentials.UnimplementedOAuthServiceServer

	listOAuthUseCase                 *usecase.ListOAuthUseCase
	getCredentialsUseCase            *usecase.GetCredentialsUseCase
	exchangeAuthorizationCodeUseCase *usecase.ExchangeAuthorizationCodeUseCase
}

func NewService(
	listOAuthUseCase *usecase.ListOAuthUseCase,
	getCredentialsUseCase *usecase.GetCredentialsUseCase,
	exchangeAuthorizationCodeUseCase *usecase.ExchangeAuthorizationCodeUseCase,
) *Service {

	service := Service{
		listOAuthUseCase:                 listOAuthUseCase,
		getCredentialsUseCase:            getCredentialsUseCase,
		exchangeAuthorizationCodeUseCase: exchangeAuthorizationCodeUseCase,
	}

	return &service
}

func (s Service) ListOAuths(request *oauthcredentials.ListOAuthsRequest, server oauthcredentials.OAuthService_ListOAuthsServer) error {
	oauthApps, err := s.listOAuthUseCase.Execute(server.Context(), request.GetOwner())
	if err != nil {
		return err
	}

	for _, oauthApp := range oauthApps {

		if err := server.Send(&oauthcredentials.ListOAuthsResponse{Oauths: []*oauthcredentials.OAuth{
			{
				Id:       oauthApp.ID,
				Owner:    oauthApp.OwnerID,
				Provider: oauthApp.Provider,
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
	//TODO implement me
	panic("implement me")
}

func (s Service) GetOAuthCredentialByProvider(ctx context.Context, request *oauthcredentials.GetOAuthCredentialByProviderRequest) (*oauthcredentials.GetOAuthCredentialByProviderResponse, error) {
	credential, err := s.getCredentialsUseCase.Execute(ctx, request.GetProvider(), request.GetOwner())
	if err != nil {
		return nil, err
	}

	return &oauthcredentials.GetOAuthCredentialByProviderResponse{
		AccessToken: credential.AccessToken,
	}, nil
}

func (s Service) ExchangeCodeForToken(ctx context.Context, request *oauthcredentials.ExchangeCodeForTokenRequest) (*oauthcredentials.ExchangeCodeForTokenResponse, error) {
	if err := s.exchangeAuthorizationCodeUseCase.Execute(ctx, request.GetProvider(), request.GetOwner(), request.GetCode()); err != nil {
		return nil, err
	}

	return &oauthcredentials.ExchangeCodeForTokenResponse{}, nil
}
