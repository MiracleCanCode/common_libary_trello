package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

const loggerKey = "logger"
const requestIDKey = "request_id"

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func GetLogger(ctx context.Context) *zap.Logger {
	return ctx.Value(loggerKey).(*zap.Logger)
}

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

func NewLoggerWithRequestID(ctx context.Context) *zap.Logger {
	logger, _ := zap.NewProduction()

	return logger.WithOptions(zap.AddCallerSkip(1), zap.Fields(zap.String("request_id", GetRequestID(ctx))))
}

func GetRequestID(ctx context.Context) string {
	requestID := ctx.Value(requestIDKey)
	if requestID == nil {
		return "unknown"
	}
	return fmt.Sprintf("%v", requestID)
}
