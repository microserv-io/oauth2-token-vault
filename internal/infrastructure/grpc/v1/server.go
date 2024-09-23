package v1

import (
	"github.com/microserv-io/oauth/internal/infrastructure/grpc/v1/oauthservice"
	"github.com/microserv-io/oauth/internal/infrastructure/grpc/v1/providerservice"
	"github.com/microserv-io/oauth/internal/usecase"
	v1 "github.com/microserv-io/oauth/pkg/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func NewServer(
	listOAuthUseCase *usecase.ListOAuthUseCase,
	getCredentialsUseCase *usecase.GetCredentialsUseCase,
) *grpc.Server {
	server := grpc.NewServer()

	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	oauthService := oauthservice.NewService(listOAuthUseCase, getCredentialsUseCase)
	providerService := providerservice.NewService()

	v1.RegisterOAuthServiceServer(server, oauthService)
	v1.RegisterOAuthProviderServiceServer(server, providerService)

	reflection.Register(server)

	return server
}
