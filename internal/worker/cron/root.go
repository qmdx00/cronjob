package cron

import (
	"github.com/robfig/cron"
	"go.uber.org/zap"
)

// RootCron implement cron.Job ...
type RootCron struct {
	log     *zap.Logger
	cronMap map[string]*cron.Cron
	keyMap  map[string]string
}

// NewRootCron ...
func NewRootCron(log *zap.Logger) cron.Job {
	return &RootCron{log: log, cronMap: make(map[string]*cron.Cron), keyMap: make(map[string]string)}
}

// Run ...
func (m *RootCron) Run() {
	// TODO
}
