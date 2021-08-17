package server

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/task/biz"
	"github.com/qmdx00/crobjob/pkg/transport"
	"github.com/qmdx00/crobjob/rpc"
	"google.golang.org/grpc"
)

var ProviderSet = wire.NewSet(NewServer, NewGRPCServer)

func NewGRPCServer(task *biz.TaskBusiness) (*grpc.Server, error) {
	// HACK: to be replaced by config
	tracer, closer, err := transport.NewJaegerTracer("cronjob_task", "127.0.0.1:6831")
	if err != nil {
		return nil, err
	}

	defer closer.Close()

	// add jaeger option
	server := grpc.NewServer(transport.JaegerServerOption(tracer))

	// HACK: register grpc service servers
	rpc.RegisterTaskServiceServer(server, task)

	return server, nil
}
