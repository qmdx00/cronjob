package log

import (
	"fmt"
	"os"
)

// Helper is a logger helper.
type Helper struct {
	logger Logger
}

// NewHelper new a logger helper.
func NewHelper(logger Logger) *Helper {
	return &Helper{
		logger: logger,
	}
}

// Log Print log by level and keyvals.
func (h *Helper) Log(level Level, keyvals ...interface{}) {
	h.logger.Log(level, keyvals...)
}

// Debug logs a message at debug level.
func (h *Helper) Debug(a ...interface{}) {
	h.logger.Log(LevelDebug, fmt.Sprint(a...))
}

// Debugf logs a message at debug level.
func (h *Helper) Debugf(format string, a ...interface{}) {
	h.logger.Log(LevelDebug, fmt.Sprintf(format, a...))
}

// Info logs a message at info level.
func (h *Helper) Info(a ...interface{}) {
	h.logger.Log(LevelInfo, fmt.Sprint(a...))
}

// Infof logs a message at info level.
func (h *Helper) Infof(format string, a ...interface{}) {
	h.logger.Log(LevelInfo, fmt.Sprintf(format, a...))
}

// Warn logs a message at warn level.
func (h *Helper) Warn(a ...interface{}) {
	h.logger.Log(LevelWarn, fmt.Sprint(a...))
}

// Warnf logs a message at warnf level.
func (h *Helper) Warnf(format string, a ...interface{}) {
	h.logger.Log(LevelWarn, fmt.Sprintf(format, a...))
}

// Error logs a message at error level.
func (h *Helper) Error(a ...interface{}) {
	h.logger.Log(LevelError, fmt.Sprint(a...))
}

// Errorf logs a message at error level.
func (h *Helper) Errorf(format string, a ...interface{}) {
	h.logger.Log(LevelError, fmt.Sprintf(format, a...))
}

// Fatal logs a message at fatal level.
func (h *Helper) Fatal(a ...interface{}) {
	h.logger.Log(LevelFatal, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatalf logs a message at fatal level.
func (h *Helper) Fatalf(format string, a ...interface{}) {
	h.logger.Log(LevelFatal, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Fatalw logs a message at fatal level.
func (h *Helper) Fatalw(keyvals ...interface{}) {
	h.logger.Log(LevelFatal, keyvals...)
	os.Exit(1)
}
