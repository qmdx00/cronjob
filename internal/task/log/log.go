package log

import (
	"github.com/qmdx00/crobjob/pkg/log"
	"go.uber.org/zap"
)

type MyLogger struct {
	logger *zap.Logger
}

func NewLogger() log.Logger {
	return &MyLogger{logger: zap.NewExample()}
}

func (l *MyLogger) Log(level log.Level, keyvals ...interface{}) {
	l.logger.Sugar().Debugw(level.String(), keyvals...)
}
