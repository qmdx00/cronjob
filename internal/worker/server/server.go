package server

import (
	"context"
	"github.com/qmdx00/crobjob/internal/worker/config"
	"github.com/qmdx00/crobjob/internal/worker/job"
	"github.com/qmdx00/crobjob/pkg/transport"
	"github.com/robfig/cron"
)

// Server ...
type Server struct {
	root   *job.RootJob
	cron   *cron.Cron
	config *config.WorkerConfig
}

// NewServer ...
func NewServer(root *job.RootJob, config *config.WorkerConfig) transport.Server {
	return &Server{root: root, cron: cron.New(), config: config}
}

// Start ...
func (s *Server) Start(_ context.Context) error {
	spec := s.config.Viper.GetString("worker.cron.spec")
	schedule, err := cron.Parse(spec)
	if err != nil {
		return err
	}

	s.cron.Schedule(schedule, s.root)
	s.cron.Run()

	return nil
}

// Stop ...
func (s *Server) Stop(_ context.Context) error {
	s.cron.Stop()
	return nil
}
