package biz

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/pkg/transport"
	"github.com/qmdx00/crobjob/rpc"
	"google.golang.org/grpc"
)

// ProviderSet for business ...
var ProviderSet = wire.NewSet(NewTaskBusiness, NewGRPCConn, NewTaskServiceClient)

func NewGRPCConn() (*grpc.ClientConn, error) {
	tracer, closer, err := transport.NewJaegerTracer("cronjob_manager", "127.0.0.1:6831")
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure(), transport.JaegerClientOption(tracer))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewTaskServiceClient(conn *grpc.ClientConn) rpc.TaskServiceClient {
	return rpc.NewTaskServiceClient(conn)
}
