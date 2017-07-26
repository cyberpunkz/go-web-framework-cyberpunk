package server

import (
	"net/http"
	"github.com/cyberpunkz/cyberpunk/logger"
)

func StartHttpServer(port string) error  {
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		logger.Error(err)
	}

	return err
}