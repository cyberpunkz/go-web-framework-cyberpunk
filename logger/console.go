package logger

import (
	"log"
	"sync"
	"fmt"
)

type Console struct{
	logConfig LogConfig
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
}

func (*Console) Error(err ...interface{}) {
	log.Printf(getPrefix() + fmt.Sprintln(err))
}

func (*Console) Debug(msg string) {
	log.Printf(getPrefix() + msg)
}
