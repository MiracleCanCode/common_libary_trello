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
	_ = logger.SetRequestID(ctx, requestID)
	log := logger.NewLoggerWithRequestID(ctx)

	log.Info("request",
		zap.String("method", info.FullMethod),
		zap.Time("request_time", time.Now()),
	)

	resStartTime := time.Now()
	res, err := handler(ctx, req)
	if err != nil {
		log.Error("Failed create response",
			zap.String("method", info.FullMethod),
		)
	}

	log.Info("response",
		zap.String("method", info.FullMethod),
		zap.Duration("response_time", time.Since(resStartTime)),
	)

	return res, err
}
