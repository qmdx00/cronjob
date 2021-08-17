package log

import (
	"go.uber.org/zap"
)

// loggerMap store loggers with different service.
var loggerMap = make(map[string]*zap.Logger)

// NewLogger get zap.Logger instance ...
func NewLogger(service string) *zap.Logger {
	if loggerMap[service] == nil {
		loggerMap[service] = zap.NewExample(zap.Fields(
			zap.String("service", service)))
	}
	return loggerMap[service]
}
