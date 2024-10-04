package grpc

import (
	"context"
	"github.com/microserv-io/oauth-credentials-server/internal/logging"
	"google.golang.org/grpc"
	"log/slog"
)

func UnaryServerInterceptor(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		ctx = logging.WithLogger(ctx, logger)
		return handler(ctx, req)
	}
}
