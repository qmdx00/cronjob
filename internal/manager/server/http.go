package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qmdx00/crobjob/internal/manager/config"
	"github.com/qmdx00/crobjob/pkg/lifecycle"
	"github.com/qmdx00/crobjob/pkg/transport"
	"go.uber.org/zap"
	"net/http"
)

// Server http server ...
type Server struct {
	log    *zap.Logger
	server *http.Server
}

// NewHttpServer ...
func NewHttpServer(log *zap.Logger, engine *gin.Engine, config *config.ManagerConfig) transport.Server {
	return &Server{
		server: &http.Server{
			Handler: engine,
			Addr:    config.Viper.GetString("manager.server.http.addr"),
		},
		log: log,
	}
}

// Start server ...
func (s *Server) Start(ctx context.Context) error {
	defer s.log.Sync()
	info, _ := lifecycle.FromContext(ctx)

	s.log.Info("server start",
		zap.String("id", info.ID()),
		zap.String("name", info.Name()),
		zap.String("version", info.Version()),
		zap.String("addr", s.server.Addr),
	)

	return s.server.ListenAndServe()
}

// Stop server ...
func (s *Server) Stop(ctx context.Context) error {
	defer s.log.Sync()
	info, _ := lifecycle.FromContext(ctx)

	s.log.Info("server stop",
		zap.String("id", info.ID()),
		zap.String("name", info.Name()),
		zap.String("version", info.Version()),
		zap.String("addr", s.server.Addr),
	)

	return s.server.Shutdown(ctx)
}
