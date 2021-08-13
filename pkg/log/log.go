package log

// Logger ...
type Logger interface {
	Log(level Level, keyvals ...interface{})
}
