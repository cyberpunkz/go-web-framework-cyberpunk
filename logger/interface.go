package logger

type Logger interface {
	AddLogConfig(logConfig LogConfig)
	Error(err error)
	Warning(msg string)
	Notice(msg string)
	Ok(msg string)
}
