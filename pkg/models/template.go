package models

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func LoadTemplates() *Template {
	htmlFiles, err := filepath.Glob("./public/html/*.html")
	if err != nil {
		fmt.Println("Error al buscar archivos HTML:", err)
		return nil
	}
	templateFiles, err := filepath.Glob("./public/templates/*.html")
	if err != nil {
		fmt.Println("Error al buscar archivos Templates:", err)
		return nil
	}
	htmlFiles = append(htmlFiles, templateFiles...)
	// return nil
	return &Template{
		templates: template.Must(template.ParseFiles(htmlFiles...)),
	}
}

func CreateTemplate(route string) *Template {
	return &Template{
		templates: template.Must(template.ParseGlob(route)),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if template := t.templates.Lookup(name); template != nil {
		return template.ExecuteTemplate(w, name, data)
	}
	return fmt.Errorf("Template not found: %s", name)
}

func (t *Template) SetTemplate(route string) {
	t.templates = template.Must(template.ParseGlob(route))
}

func (t *Template) GetTemplate() *template.Template {
	return t.templates
}
