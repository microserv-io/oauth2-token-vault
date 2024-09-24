package v1

import (
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/grpc/v1/oauthservice"
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/grpc/v1/providerservice"
	"github.com/microserv-io/oauth-credentials-server/internal/usecase"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
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

	oauthcredentials.RegisterOAuthServiceServer(server, oauthService)
	oauthcredentials.RegisterOAuthProviderServiceServer(server, providerService)

	reflection.Register(server)

	return server
}
