package server

import (
	"github.com/robfig/cron"
	"go.uber.org/zap"
)

// MainCron implement cron.Job ...
type MainCron struct {
	log *zap.Logger
}

// NewMainCron ...
func NewMainCron(log *zap.Logger) cron.Job {
	return &MainCron{log: log}
}

// Run ...
func (m *MainCron) Run() {
	// TODO
}
