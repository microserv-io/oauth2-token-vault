package logging

import (
	"context"
	"log/slog"
	"os"
)

type contextKey struct{}

var loggerKey = &contextKey{}

var _ = LoggerFromContext

func LoggerFromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	if !ok {
		return slog.New(slog.NewTextHandler(os.Stdout, nil)) // Default logger
	}
	return logger
}

func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}
