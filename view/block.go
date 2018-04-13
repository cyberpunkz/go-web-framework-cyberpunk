package view

import (
	"fmt"
)

type BlockContainer struct {
	templates map[string]string
	data      map[string]interface{}
}

func NewBlockContainer() *BlockContainer {
	return &BlockContainer{
		templates: make(map[string]string),
		data:      make(map[string]interface{}),
	}
}

func (r *BlockContainer) AddBlock(name string, template string, data interface{}) error {
	if _, ok := r.templates[name]; ok {
		return fmt.Errorf("block %s already exists", name)

	} else {
		r.templates[name] = template
		r.data[name] = data

		return nil
	}
}

func (r *BlockContainer) UpdateBlock(name string, template string, data interface{}) error {
	if _, ok := r.templates[name]; ok {
		r.templates[name] = template
		r.data[name] = data

		return nil
	} else {
		return fmt.Errorf("block %s does not exist", name)
	}
}
