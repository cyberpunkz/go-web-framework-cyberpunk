package logger

import (
	"log"
	"sync"
	"os"
	"fmt"
)

type File struct {
	logConfig LogConfig
	error     *log.Logger
	debug     *log.Logger
}

var file *File
var fileOnce sync.Once

func GetFile() *File {
	fileOnce.Do(func() {
		file = &File{}
	})

	return file
}

func (*File) AddLogConfig(logConfig LogConfig) {
	file.logConfig = logConfig
	flag := log.LstdFlags | log.Lmicroseconds

	out, _ := os.Create(logConfig.Path + "/error.log")
	file.error = log.New(out, "Error: ", flag)

	out, _ = os.Create(logConfig.Path + "/debug.log")
	file.debug = log.New(out, "Debug: ", flag)
}

func (*File) Error(err ...interface{}) {
	msg := getPrefix()
	file.error.Printf(fmt.Sprintln(err) + msg)
	log.Printf(fmt.Sprintln(err) + msg)
}

func (*File) Debug(msg string) {
	msg = getPrefix() + msg
	file.debug.Printf(msg)
	log.Printf(msg)
}
