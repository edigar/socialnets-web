package template

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

func ExecuteTemplate(w http.ResponseWriter, template string, data any) {
	err := templates.ExecuteTemplate(w, template, data)
	if err != nil {
		fmt.Println(err)
	}
}
