package view

import (
	"net/http"
	"html/template"
	"errors"
	"github.com/cyberpunkz/cyberpunk/proxy"
)

func Render(w http.ResponseWriter, routeName string, page interface{}) (err error) {
	templateRegister := GetTemplateRegister()
	templateFiles := templateRegister.getTemplates(routeName)

	if len(templateFiles) == 0 {
		err = errors.New("There are no registered templates to rendering the page!")
		proxy.Error().Println(err)
		return
	}

	templates := template.Must(template.ParseFiles(templateFiles...))
	err = templates.ExecuteTemplate(w, routeName + ".html", page)

	if err != nil {
		proxy.Error().Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}
