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

// Server grpc server ...
type Server struct {
	log    *zap.Logger
	server *grpc.Server
}

// NewServer ...
func NewServer(log *zap.Logger, server *grpc.Server) transport.Server {
	return &Server{log: log, server: server}
}

// Start server ...
func (s *Server) Start(ctx context.Context) error {
	defer s.log.Sync()
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

	return s.server.Serve(lis)

}

// Stop server ...
func (s *Server) Stop(ctx context.Context) error {
	defer s.log.Sync()
	info, _ := lifecycle.FromContext(ctx)

	s.log.Info("server stop",
		zap.String("id", info.ID()),
		zap.String("name", info.Name()),
		zap.String("version", info.Version()),
		zap.String("addr", info.Metadata()[constant.GRPCAddr]),
	)

	s.server.GracefulStop()
	return nil
}
