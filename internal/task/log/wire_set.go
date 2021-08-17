package log

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/pkg/log"
	"go.uber.org/zap"
)

// ProviderSet for log ...
var ProviderSet = wire.NewSet(NewTaskLogger)

// NewTaskLogger ...
func NewTaskLogger() *zap.Logger {
	return log.NewLogger("task-service")
}
