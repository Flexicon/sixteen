package views

import (
	"embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

//go:embed *.html
var embeddedTemplates embed.FS

type renderer struct {
	templates *template.Template
}

// NewRenderer constructor
func NewRenderer() echo.Renderer {
	return &renderer{
		templates: template.Must(template.ParseFS(embeddedTemplates, "*")),
	}
}

func (t *renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
