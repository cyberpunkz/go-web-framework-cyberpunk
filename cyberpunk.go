package cyberpunk

import (
	"github.com/cyberpunkz/cyberpunk/server"
	"github.com/cyberpunkz/cyberpunk/config/system"
	"github.com/cyberpunkz/cyberpunk/view"
	"net/http"
	"github.com/cyberpunkz/cyberpunk/routes"
	"log"
)

func GetSystemConfig() *system.Config {
	return system.GetInstance()
}

func Error() *log.Logger {
	return GetSystemConfig().GetCurrentLogger().Error()
}

func Debug() *log.Logger {
	return GetSystemConfig().GetCurrentLogger().Debug()
}

func GetTemplateRegister() *view.TemplateRegister {
	return view.GetTemplateRegister()
}

func AddNewRoute(route string, handler func(http.ResponseWriter, *http.Request)) {
	routes.AddNewRoute(route, handler)
}

func RenderLayout(w http.ResponseWriter, routeName string, page interface{}) error {
	return view.Render(w, routeName, page)
}

func StartNewHttpServer(port string) error {
	return server.StartHttpServer(port)
}
