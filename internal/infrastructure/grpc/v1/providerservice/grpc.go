package providerservice

import (
	"context"
	grpcoauthv1 "github.com/microserv-io/oauth/pkg/grpc/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	grpcoauthv1.UnimplementedOAuthProviderServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s Service) ListProviders(_ *emptypb.Empty, stream grpcoauthv1.OAuthProviderService_ListProvidersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListProviders not implemented")
}
func (s Service) CreateProvider(ctx context.Context, oauthProvider *grpcoauthv1.CreateProviderRequest) (*grpcoauthv1.OAuthProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProvider not implemented")
}
func (s Service) UpdateProvider(ctx context.Context, oauthProvider *grpcoauthv1.UpdateProviderRequest) (*grpcoauthv1.OAuthProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProvider not implemented")
}
func (s Service) DeleteProvider(ctx context.Context, oauthProvider *grpcoauthv1.DeleteProviderRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProvider not implemented")
}
