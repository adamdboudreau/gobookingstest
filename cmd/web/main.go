package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"bookings/pkg/models"
	"bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

// mac run bunch of go files together "go run *.go"
// windows run bunch of go files together "go run ."
const portNumber = ":3030"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("starting app on port %s", portNumber))
	// http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	gob.Register(models.Reservation{}) // allow storing reservation model in session

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	app.UseCache = false
	tc, err := render.CreateTemplateCache()
	if err != nil {
		// log.Fatal("cannot load templates")
		return err
	}
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world !<div>ok</div>")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})*/

	return nil
}
