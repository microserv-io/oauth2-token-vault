package grpc

import (
	"github.com/microserv-io/oauth-credentials-server/internal/app/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/app/provider"
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/grpc/oauthservice/v1"
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/grpc/providerservice/v1"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log/slog"
)

func NewServer(
	oauthAppService *oauthapp.Service,
	providerService *provider.Service,
	logger *slog.Logger,
) *grpc.Server {

	server := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerInterceptor(logger)))

	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	oauthService := v1.NewService(oauthAppService)
	providerServiceGrpc := providerservice.NewService(providerService, oauthAppService)

	oauthcredentials.RegisterOAuthServiceServer(server, oauthService)
	oauthcredentials.RegisterOAuthProviderServiceServer(server, providerServiceGrpc)

	reflection.Register(server)

	return server
}
