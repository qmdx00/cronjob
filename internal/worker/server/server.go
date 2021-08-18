package server

import (
	"context"
	"github.com/qmdx00/crobjob/pkg/transport"
	"github.com/robfig/cron"
)

// Server ...
type Server struct {
	job   cron.Job
	cron  *cron.Cron
}

// NewServer ...
func NewServer(job cron.Job) transport.Server {
	return &Server{job: job, cron: cron.New()}
}

// Start ...
func (s *Server) Start(ctx context.Context) error {
	// HACK: replace spec from config
	schedule, err := cron.Parse("0/1 * * * * ?")
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
