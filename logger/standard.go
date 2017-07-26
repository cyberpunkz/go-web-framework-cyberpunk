package logger

import "log"

func Error(err error) {
	log.Panic(err)
}
