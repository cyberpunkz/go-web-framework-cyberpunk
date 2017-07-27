package logger

import (
	"log"
	"sync"
	"os"
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
	flag := log.LstdFlags | log.Lmicroseconds | log.Llongfile

	out, _ := os.Create(logConfig.Path + "/error.log")
	file.error = log.New(out, "Error: ", flag)

	out, _ = os.Create(logConfig.Path + "/debug.log")
	file.debug = log.New(out, "Debug: ", flag)
}

func (*File) Error() *log.Logger {
	return file.error
}

func (*File) Debug() *log.Logger {
	return file.debug
}
