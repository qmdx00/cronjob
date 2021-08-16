package biz

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/task/constant"
	"github.com/qmdx00/crobjob/rpc"
	"google.golang.org/grpc"
)

// ProviderSet for business ...
var ProviderSet = wire.NewSet(NewTaskBusiness, NewGRPCConn, NewTaskServiceClient)

func NewGRPCConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(constant.GRPCAddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewTaskServiceClient(conn *grpc.ClientConn) rpc.TaskServiceClient {
	return rpc.NewTaskServiceClient(conn)
}
