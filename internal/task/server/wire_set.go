package server

import (
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"github.com/qmdx00/crobjob/internal/task/biz"
	"github.com/qmdx00/crobjob/pkg/middleware"
	"github.com/qmdx00/crobjob/rpc"
	"google.golang.org/grpc"
)

// ProviderSet for server ...
var ProviderSet = wire.NewSet(NewServer, NewGRPCServer)

// NewGRPCServer ...
func NewGRPCServer(task *biz.TaskBusiness, tracer opentracing.Tracer) (*grpc.Server, error) {
	// add jaeger option
	server := grpc.NewServer(middleware.JaegerServerOption(tracer))

	// HACK: register grpc service servers
	rpc.RegisterTaskServiceServer(server, task)

	return server, nil
}
