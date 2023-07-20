package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"bookings/pkg/helpers"
	"bookings/pkg/models"
	"bookings/pkg/render"
	"bookings/pkg/storage/mysql"

	"github.com/alexedwards/scs/v2"
)

// mac run bunch of go files together "go run *.go"
// windows run bunch of go files together "go run ."
const portNumber = ":3030"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

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

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

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

	mysqlDB, err := mysql.NewStorageRepository("", "", "")
	if err != nil {
		// log.Fatal("cannot load templates")
		return err
	}
	app.DB = mysqlDB
	// test local query
	// testRes := models.Reservation{FirstName: "test", LastName: "user", Email: "test@abc.com", Phone: "1234"}
	// ctx := context.Background() // background nil context
	// app.DB.SaveReservation(ctx, testRes)
	// res, err := app.DB.FindReservation(ctx, testRes.Email)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("found user ", res.FirstName, res.LastName)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

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
