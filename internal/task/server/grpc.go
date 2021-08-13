package server

import (
	"context"
	log "github.com/qmdx00/crobjob/pkg/log"
	"github.com/qmdx00/crobjob/pkg/transport"
)

type GRPCServer struct {
	helper *log.Helper
}

func NewGRPCServer(helper *log.Helper) (transport.Server, func(), error) {
	return &GRPCServer{helper: helper}, func() {
		helper.Error("cleanup")
	}, nil
}

func (s *GRPCServer) Start(ctx context.Context) error {
	s.helper.Debug("server start")
	return nil
}

func (s *GRPCServer) Stop(ctx context.Context) error {
	s.helper.Debug("server stop")
	return nil
}
