package oauthservice

import (
	"context"
	"github.com/microserv-io/oauth/internal/usecase"
	v1 "github.com/microserv-io/oauth/pkg/grpc/v1"
)

type Service struct {
	v1.UnimplementedOAuthServiceServer

	listOAuthUseCase      *usecase.ListOAuthUseCase
	getCredentialsUseCase *usecase.GetCredentialsUseCase
}

func NewService(
	listOAuthUseCase *usecase.ListOAuthUseCase,
	getCredentialsUseCase *usecase.GetCredentialsUseCase,
) *Service {

	service := Service{
		listOAuthUseCase:      listOAuthUseCase,
		getCredentialsUseCase: getCredentialsUseCase,
	}

	return &service
}

func (s Service) ListOAuths(request *v1.ListOAuthsRequest, server v1.OAuthService_ListOAuthsServer) error {
	oauthApps, err := s.listOAuthUseCase.Execute(server.Context(), request.GetOwner())
	if err != nil {
		return err
	}

	for _, oauthApp := range oauthApps {
		if err := server.Send(&v1.OAuth{
			Id:       oauthApp.ID,
			Owner:    oauthApp.OwnerID,
			Provider: oauthApp.Provider,
			Scopes:   oauthApp.Scopes,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (s Service) GetOAuthByID(ctx context.Context, request *v1.GetOAuthByIDRequest) (*v1.OAuth, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetOAuthByProvider(ctx context.Context, request *v1.GetOAuthByProviderRequest) (*v1.OAuth, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetOAuthCredentialsByProvider(ctx context.Context, request *v1.GetOAuthByProviderRequest) (*v1.OAuthCredential, error) {
	credential, err := s.getCredentialsUseCase.Execute(ctx, request.GetProvider(), request.GetOwner())
	if err != nil {
		return nil, err
	}

	return &v1.OAuthCredential{
		AccessToken: credential.AccessToken,
	}, nil
}
