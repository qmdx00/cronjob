package server

import (
	"github.com/robfig/cron"
)

// MainCron implement cron.Job ...
type MainCron struct {
}

// NewMainCron ...
func NewMainCron() cron.Job {
	return &MainCron{}
}

// Run ...
func (m *MainCron) Run() {
	// TODO
}
