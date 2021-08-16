package log

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/pkg/log"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(NewTaskLogger)

func NewTaskLogger() *zap.Logger {
	return log.NewLogger("task-service")
}
