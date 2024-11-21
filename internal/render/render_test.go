package render

import (
	"net/http"
	"testing"

	"github.com/Badi96/go-bookings-website/internal/models"
)

func TestAddDefaultData(t *testing.T) {

	var td models.TemplateData
	r, err := getSession()

	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDataDefault(&td, r)

	if result.Flash != "123" {
		t.Error("Flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	var ww myWriter
	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser", err)
	}

	err = RenderTemplate(&ww, r, "template-that-does-not-exist.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template thsat does not existr")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session")) // add session data in context
	r = r.WithContext(ctx)
	return r, nil
}

func TestNewTamplates(t *testing.T) {
	NewTamplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
