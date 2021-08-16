package server

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/task/biz"
	"github.com/qmdx00/crobjob/rpc"
	"google.golang.org/grpc"
)

var ProviderSet = wire.NewSet(NewServer, NewGRPC)

func NewGRPC(task *biz.TaskBusiness) *grpc.Server {
	server := grpc.NewServer()

	// HACK: register grpc service servers
	rpc.RegisterTaskServiceServer(server, task)

	return server
}
