package log

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/task/config"
	"github.com/qmdx00/crobjob/pkg/log"
	"go.uber.org/zap"
)

// ProviderSet for log ...
var ProviderSet = wire.NewSet(NewTaskLogger)

// NewTaskLogger ...
func NewTaskLogger(config *config.TaskConfig) *zap.Logger {
	serviceName := config.Viper.GetString("task.log.prefix")
	return log.NewLogger(serviceName)
}
