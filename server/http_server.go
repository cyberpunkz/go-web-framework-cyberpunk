package server

import (
	"net/http"
	"github.com/cyberpunkz/cyberpunk/proxy"
)

func StartHttpServer(port string) error  {
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		proxy.Log().Error(err)
	}

	return err
}