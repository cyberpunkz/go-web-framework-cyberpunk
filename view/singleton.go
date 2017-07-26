package view

import (
	"sync"
	"github.com/cyberpunkz/cyberpunk/config/system"
	"errors"
)

type Routes struct {
	name      string
	templates []string
}

type TemplateRegister struct {
	routes []Routes
}

var config = system.GetInstance()
var instance *TemplateRegister
var once sync.Once

func GetTemplateRegister() *TemplateRegister {
	once.Do(func() {
		instance = &TemplateRegister{}
	})
	return instance
}

func (instance *TemplateRegister) Register(routeName string) {
	instance.routes = append(
		instance.routes,
		Routes{name:routeName})
}

func (instance *TemplateRegister) AddTemplate(routeName string, template string) (err error) {
	found := false

	for key, route := range instance.routes {
		if route.name == routeName {
			instance.routes[key].templates = append(
				instance.routes[key].templates,
				config.GetFrontendDir() + "/html/" + template + ".html")
			found = true; break
		}
	}

	if found != true {
		err = errors.New("The given routName: '" + routeName + "' does not registered!")
	}

	return
}

func (instance *TemplateRegister) getTemplates(routeName string) (files []string) {
	for key, route := range instance.routes {
		if route.name == routeName {
			files = instance.routes[key].templates
			break
		}
	}

	return files
}
