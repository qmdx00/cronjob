package server

import (
	"context"
	"github.com/qmdx00/crobjob/internal/worker/config"
	"github.com/qmdx00/crobjob/pkg/transport"
	"github.com/robfig/cron"
)

// Server ...
type Server struct {
	job    cron.Job
	cron   *cron.Cron
	config *config.WorkerConfig
}

// NewServer ...
func NewServer(job cron.Job, config *config.WorkerConfig) transport.Server {
	return &Server{job: job, cron: cron.New(), config: config}
}

// Start ...
func (s *Server) Start(ctx context.Context) error {
	spec := s.config.Viper.GetString("worker.cron.spec")
	schedule, err := cron.Parse(spec)
	if err != nil {
		return err
	}

	s.cron.Schedule(schedule, s.job)
	s.cron.Run()

	return nil
}

// Stop ...
func (s *Server) Stop(ctx context.Context) error {
	s.cron.Stop()
	return nil
}
