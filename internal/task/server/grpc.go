package server

import (
	"context"
	"fmt"
	"github.com/qmdx00/crobjob/pkg/transport"
	"log"
)

type GRPCServer struct {
}

func NewGRPCServer() (transport.Server, func(), error) {
	return &GRPCServer{}, func() {
		fmt.Println("cleanup")
	}, nil
}

func (s *GRPCServer) Start(ctx context.Context) error {
	log.Println("server start")
	return nil
}

func (s *GRPCServer) Stop(ctx context.Context) error {
	log.Println("server stop")
	return nil
}
