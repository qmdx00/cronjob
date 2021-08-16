package server

import (
	"context"
	"github.com/qmdx00/crobjob/internal/task/constant"
	"github.com/qmdx00/crobjob/pkg/lifecycle"
	"github.com/qmdx00/crobjob/pkg/transport"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	log  *zap.Logger
	grpc *grpc.Server
}

func NewServer(log *zap.Logger, grpc *grpc.Server) (transport.Server, error) {
	return &Server{log: log, grpc: grpc}, nil
}

func (s *Server) Start(ctx context.Context) error {
	info, _ := lifecycle.FromContext(ctx)

	lis, err := net.Listen("tcp", info.Metadata()[constant.GRPCAddr])
	if err != nil {
		s.log.Fatal("failed to listen", zap.Error(err))
	}

	s.log.Info("server start",
		zap.String("id", info.ID()),
		zap.String("name", info.Name()),
		zap.String("version", info.Version()),
		zap.String("addr", info.Metadata()[constant.GRPCAddr]),
	)

	return s.grpc.Serve(lis)

}

func (s *Server) Stop(ctx context.Context) error {
	info, _ := lifecycle.FromContext(ctx)

	s.log.Info("server stop",
		zap.String("id", info.ID()),
		zap.String("name", info.Name()),
		zap.String("version", info.Version()),
		zap.String("addr", info.Metadata()[constant.GRPCAddr]),
	)

	s.grpc.GracefulStop()
	return nil
}
