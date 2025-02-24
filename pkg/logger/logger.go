package logger

import (
	"context"

	"go.uber.org/zap"
)

const loggerKey = "logger"

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func GetLogger(ctx context.Context) *zap.Logger {
	return ctx.Value(loggerKey).(*zap.Logger)
}
