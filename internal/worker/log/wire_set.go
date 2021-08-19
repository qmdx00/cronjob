package log

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/worker/config"
	"github.com/qmdx00/crobjob/pkg/log"
	"go.uber.org/zap"
)

// ProviderSet for log ...
var ProviderSet = wire.NewSet(NewWorkerLogger)

// NewWorkerLogger ...
func NewWorkerLogger(config *config.WorkerConfig) *zap.Logger {
	serviceName := config.Viper.GetString("worker.log.prefix")
	return log.NewLogger(serviceName)
}
