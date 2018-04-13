package view

import (
	"html/template"
	"io"
)

type Renderer struct {
	dir  string
	tmpl *template.Template // it should stored in cache not here
}

func NewRenderer(dir string) *Renderer {
	return &Renderer{
		dir: dir,
	}
}

func (r *Renderer) Parse(registry *BlockContainer) {
	var templates []string

	for _, templateName := range registry.templates {
		templates = append(templates, r.dir+templateName)
	}

	r.tmpl = template.Must(template.ParseFiles(templates...))
}

func (r *Renderer) Render(wr io.Writer, name string, registry *BlockContainer) error {
	return r.tmpl.ExecuteTemplate(wr, name, registry.data)
}
