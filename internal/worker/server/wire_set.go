package server

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/worker/config"
	"github.com/qmdx00/crobjob/internal/worker/consumer"
	"github.com/qmdx00/crobjob/pkg/transport"
	"github.com/robfig/cron"
)

// ProviderSet for cron job ...
var ProviderSet = wire.NewSet(
	NewMainCron,
	NewServers,
)

func NewServers(job cron.Job, config *config.WorkerConfig) ([]transport.Server, error) {
	servers := make([]transport.Server, 0)

	// add cron task server
	cronServer := NewServer(job, config)
	servers = append(servers, cronServer)

	// add kafka consumer client
	taskServer, err := consumer.NewTaskConsumer(config)
	if err != nil {
		return nil, err
	}
	servers = append(servers, taskServer)

	return servers, nil
}
