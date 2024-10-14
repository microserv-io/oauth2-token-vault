package grpc

import (
	"github.com/microserv-io/oauth2-token-vault/internal/app/oauthapp"
	"github.com/microserv-io/oauth2-token-vault/internal/app/provider"
	"github.com/microserv-io/oauth2-token-vault/internal/infrastructure/grpc/v1"
	"github.com/microserv-io/oauth2-token-vault/pkg/proto/oauthcredentials/v1"
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

	oauthServiceGrpc := v1.NewOAuthAppServiceGRPC(oauthAppService)
	providerServiceGrpc := v1.NewProviderServiceGRPC(providerService)

	oauthcredentials.RegisterOAuthServiceServer(server, oauthServiceGrpc)
	oauthcredentials.RegisterOAuthProviderServiceServer(server, providerServiceGrpc)

	reflection.Register(server)

	return server
}
