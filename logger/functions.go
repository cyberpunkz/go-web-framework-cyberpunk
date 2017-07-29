package logger

import (
	"runtime"
	"strconv"
)

func getPrefix() string {
	_, filePath, line, ok := runtime.Caller(2)

	if !ok {
		filePath = "???"
		line = 0
	}

	return filePath + ":" + strconv.Itoa(line) +": "
}
