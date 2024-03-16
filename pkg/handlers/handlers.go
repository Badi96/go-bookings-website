package handlers

import (
	"net/http"

	"github.com/Badi96/Golang-Bed-Breakfast-booking-website/pkg/render"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "about.page.tmpl")
}
