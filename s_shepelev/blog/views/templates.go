package views

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// TemplateRegistry - define the template registry struct
type TemplateRegistry struct {
	Templates map[string]*template.Template
}

// Render - implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.Templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}

	return tmpl.ExecuteTemplate(w, "layout", data)
}
