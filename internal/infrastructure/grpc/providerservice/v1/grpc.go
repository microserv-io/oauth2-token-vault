package v1

import (
	"context"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ oauthcredentials.OAuthProviderServiceServer = &Service{}

type Service struct {
	oauthcredentials.UnimplementedOAuthProviderServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s Service) ListProviders(_ *oauthcredentials.ListProvidersRequest, stream oauthcredentials.OAuthProviderService_ListProvidersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListProviders not implemented")
}
func (s Service) CreateProvider(ctx context.Context, oauthProvider *oauthcredentials.CreateProviderRequest) (*oauthcredentials.CreateProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProvider not implemented")
}
func (s Service) UpdateProvider(ctx context.Context, oauthProvider *oauthcredentials.UpdateProviderRequest) (*oauthcredentials.UpdateProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProvider not implemented")
}
func (s Service) DeleteProvider(ctx context.Context, oauthProvider *oauthcredentials.DeleteProviderRequest) (*oauthcredentials.DeleteProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProvider not implemented")
}
