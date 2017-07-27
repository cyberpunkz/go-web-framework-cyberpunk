package logger

import (
	"log"
	"sync"
	"os"
)

type File struct{
	logConfig LogConfig
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
}

func (*File) Error(err error) {
	writeLog(file.logConfig.Path + "/error", err.Error())
}

func (*File) Warning(msg string) {
	writeLog(file.logConfig.Path + "/warning", msg)
}

func (*File) Notice(msg string) {
	writeLog(file.logConfig.Path + "/notice", msg)
}

func (*File) Ok(msg string) {
	writeLog(file.logConfig.Path + "/ok", msg)
}

func writeLog(filePath string, text string) {
	f, err := os.OpenFile(
		filePath + ".log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0600,
	)

	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text + "\r\n"); err != nil {
		log.Println(err)
	}
}
