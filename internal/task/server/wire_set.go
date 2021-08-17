package server

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/task/biz"
	"github.com/qmdx00/crobjob/pkg/middleware"
	"github.com/qmdx00/crobjob/rpc"
	"google.golang.org/grpc"
)

// ProviderSet for server ...
var ProviderSet = wire.NewSet(NewServer, NewGRPCServer)

// NewGRPCServer ...
func NewGRPCServer(task *biz.TaskBusiness) (*grpc.Server, func(), error) {
	// HACK: to be replaced by config
	tracer, closer, err := middleware.NewJaegerTracer("cronjob_task", "127.0.0.1:6831")
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
