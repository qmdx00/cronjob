package server

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/worker/config"
	"github.com/qmdx00/crobjob/internal/worker/job"
	"github.com/qmdx00/crobjob/pkg/transport"
	"go.uber.org/zap"
)

// ProviderSet for cron job ...
var ProviderSet = wire.NewSet(NewServers)

func NewServers(root *job.RootJob, config *config.WorkerConfig, log *zap.Logger, receive Receive) ([]transport.Server, error) {
	servers := make([]transport.Server, 0)

	// add root cron task
	cronServer := NewServer(root, config)
	servers = append(servers, cronServer)

	// add kafka handler client
	taskServer, err := NewTaskConsumer(config, log, receive)
	if err != nil {
		return nil, err
	}
	servers = append(servers, taskServer)

	return servers, nil
}
