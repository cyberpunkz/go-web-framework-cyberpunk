package logger

import (
	"log"
	"sync"
	"os"
)

type Console struct{
	logConfig LogConfig
	debug *log.Logger
	error *log.Logger
}

var console *Console
var consoleOnce sync.Once

func GetConsole() *Console {
	consoleOnce.Do(func() {
		console = &Console{}
	})

	return console
}

func (*Console) AddLogConfig(logConfig LogConfig) {
	console.logConfig = logConfig

	console.debug = log.New(os.Stdout, "Debug: ", log.Ldate|log.Ltime|log.Llongfile)
	console.error = log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Llongfile)
}
func (*Console) Error() *log.Logger {
	return console.error
}

func (*Console) Debug() *log.Logger {
	return console.debug
}
