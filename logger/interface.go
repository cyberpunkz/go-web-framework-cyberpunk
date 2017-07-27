package logger

import "log"

type Logger interface {
	AddLogConfig(logConfig LogConfig)
	Error() *log.Logger
	Debug() *log.Logger
}
