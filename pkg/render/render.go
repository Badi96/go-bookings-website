package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Badi96/go-bookings-website/pkg/config"
	"github.com/Badi96/go-bookings-website/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config for the package
func NewTamplates(a *config.AppConfig) {
	app = a
}

func AddDataDefault(td *models.TemplateData) *models.TemplateData {
	return td
}

// RendrTemplate redners template using html/template
// tmpl name of string we want to render
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	//Is not in production, read directly from desk, otherwise read from our template cache.
	if app.UseCache {

		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cahce")
	}
	buf := new(bytes.Buffer)

	td = AddDataDefault(td) // passing in the default data to template data
	// take template, execute it, don't pass any data, and store it in buf variable
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// CreateTemplateCache Creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	mycache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return mycache, err
	}

	for _, page := range pages {
		//returning base of the file instead of entire path
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return mycache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return mycache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(("./templates/*.layout.tmpl"))
			if err != nil {
				return mycache, err
			}
		}
		mycache[name] = ts

	}
	return mycache, nil
}
