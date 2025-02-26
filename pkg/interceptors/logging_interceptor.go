package interceptors

import (
	"context"
	"time"

	"github.com/MiracleCanCode/common_libary_trello/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LoggerInterceptor(
	ctx context.Context,
	req any, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	requestID := uuid.New().String()
	log := logger.GetLogger(ctx)

	log.Info("request", zap.String("request_id", requestID),
		zap.String("method", info.FullMethod),
		zap.Time("request_time", time.Now()),
	)

	res, err := handler(ctx, req)
	return res, err
}
