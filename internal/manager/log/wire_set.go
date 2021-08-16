package log

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/pkg/log"
	"go.uber.org/zap"
)

// ProviderSet ...
var ProviderSet = wire.NewSet(NewManagerLogger)

func NewManagerLogger() *zap.Logger {
	return log.NewLogger("manager-service")
}
