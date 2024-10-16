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
