package logger

import (
	"log"
	"sync"
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
func (*Console) Error(err error) {
	log.Println(err)
}

func (*Console) Warning(msg string) {
	log.Printf("Warning: " + msg)
}

func (*Console) Notice(msg string) {
	log.Printf("Notice: " + msg)
}

func (*Console) Ok(msg string) {
	log.Printf("Ok: " + msg)
}
