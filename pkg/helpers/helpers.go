package helpers

import (
	"bookings/pkg/config"
	"fmt"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

// NewHelpers setup app config for helpers
func NewHelpers(appConfig *config.AppConfig) {
	app = appConfig
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client Error with status of ", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
