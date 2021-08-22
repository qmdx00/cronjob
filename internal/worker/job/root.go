package job

import (
	"github.com/robfig/cron"
	"go.uber.org/zap"
)

// RootJob implement job.Job ...
type RootJob struct {
	log *zap.Logger
}

// NewRootJob ...
func NewRootJob(log *zap.Logger) cron.Job {
	return &RootJob{log: log}
}

// Run ...
func (m *RootJob) Run() {
	// TODO
}
