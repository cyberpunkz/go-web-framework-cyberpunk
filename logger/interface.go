package logger

type Logger interface {
	AddLogConfig(logConfig LogConfig)
	Error(err ...interface{})
	Debug(msg string)
}
