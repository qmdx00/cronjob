package biz

import (
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"github.com/qmdx00/crobjob/internal/manager/config"
	"github.com/qmdx00/crobjob/pkg/middleware"
	"github.com/qmdx00/crobjob/rpc"
	"google.golang.org/grpc"
)

// ProviderSet for business ...
var ProviderSet = wire.NewSet(NewGRPCConn, NewTracer, NewTaskBusiness, NewTaskServiceClient)

// NewGRPCConn ...
func NewGRPCConn(tracer opentracing.Tracer, config *config.ManagerConfig) (*grpc.ClientConn, func(), error) {
	conn, err := grpc.Dial(
		config.Viper.GetString("task.server.grpc.addr"),
		grpc.WithInsecure(),
		middleware.JaegerClientOption(tracer),
	)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() {
		_ = conn.Close()
	}, nil
}

// NewTracer ...
func NewTracer(config *config.ManagerConfig) (opentracing.Tracer, func(), error) {
	tracer, closer, err := middleware.NewJaegerTracer(
		config.Viper.GetString("manager.log.prefix"),
		config.Viper.GetString("resource.jaeger.agent"),
	)
	if err != nil {
		return nil, nil, err
	}
	return tracer, func() {
		_ = closer.Close()
	}, nil
}

// NewTaskServiceClient ...
func NewTaskServiceClient(conn *grpc.ClientConn) rpc.TaskServiceClient {
	return rpc.NewTaskServiceClient(conn)
}
