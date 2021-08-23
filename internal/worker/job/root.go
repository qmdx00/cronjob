package job

import (
	"go.uber.org/zap"
)

// RootJob implement job.Job ...
type RootJob struct {
	log *zap.Logger
}

// NewRootJob ...
func NewRootJob(log *zap.Logger) *RootJob {
	return &RootJob{log: log}
}

// Run ...
func (m *RootJob) Run() {
	// TODO
}
