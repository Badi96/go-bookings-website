package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Badi96/go-bookings-website/internal/config"
	"github.com/Badi96/go-bookings-website/internal/handlers"
	"github.com/Badi96/go-bookings-website/internal/models"
	"github.com/Badi96/go-bookings-website/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	//what am I going to put in the session?
	gob.Register(models.Reservation{})

	//change this to true, when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true //keep cookie after browser window is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = session

	// encrypt cookie?
	session.Cookie.Secure = app.InProduction

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache!", err)
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
