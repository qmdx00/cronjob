package log

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/manager/config"
	"github.com/qmdx00/crobjob/pkg/log"
	"go.uber.org/zap"
)

// ProviderSet ...
var ProviderSet = wire.NewSet(NewManagerLogger)

// NewManagerLogger ...
func NewManagerLogger(config *config.ManagerConfig) *zap.Logger {
	serviceName := config.Viper.GetString("manager.log.prefix")
	return log.NewLogger(serviceName)
}
