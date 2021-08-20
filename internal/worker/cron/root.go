package cron

import (
	"github.com/robfig/cron"
	"go.uber.org/zap"
)

// RootCron implement cron.Job ...
type RootCron struct {
	log *zap.Logger
}

// NewRootCron ...
func NewRootCron(log *zap.Logger) cron.Job {
	return &RootCron{log: log}
}

// Run ...
func (m *RootCron) Run() {
	// TODO
}
