package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Badi96/Golang-Bed-Breakfast-booking-website/pkg/config"
	"github.com/Badi96/Golang-Bed-Breakfast-booking-website/pkg/handlers"
	"github.com/Badi96/Golang-Bed-Breakfast-booking-website/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache!")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTamplates(&app)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
