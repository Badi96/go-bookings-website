package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RendrTemplate redners template using html/template
// tmpl name of string we want to render
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}
