package server

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/task/biz"
	"github.com/qmdx00/crobjob/internal/task/config"
	"github.com/qmdx00/crobjob/pkg/middleware"
	"github.com/qmdx00/crobjob/rpc"
	"google.golang.org/grpc"
)

// ProviderSet for server ...
var ProviderSet = wire.NewSet(NewServer, NewGRPCServer)

// NewGRPCServer ...
func NewGRPCServer(task *biz.TaskBusiness, config *config.TaskConfig) (*grpc.Server, func(), error) {
	serviceName := config.Viper.GetString("task.log.prefix")
	agent := config.Viper.GetString("resource.jaeger.agent")

	tracer, closer, err := middleware.NewJaegerTracer(serviceName, agent)
	if err != nil {
		return nil, nil, err
	}

	// add jaeger option
	server := grpc.NewServer(middleware.JaegerServerOption(tracer))

	// HACK: register grpc service servers
	rpc.RegisterTaskServiceServer(server, task)

	return server, func() {
		_ = closer.Close()
	}, nil
}
