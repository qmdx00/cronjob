package server

import (
	"fmt"
	"github.com/robfig/cron"
)

// MainCron ...
type MainCron struct {
}

// NewMainCron ...
func NewMainCron() cron.Job {
	return &MainCron{}
}

// Run implement cron.Job ...
func (m *MainCron) Run() {
	fmt.Println("hello world")
}
