package routes

import (
	"net/http"
)

func AddNewRoute(route string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(route, handler)
}
