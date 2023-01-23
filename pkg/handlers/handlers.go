package handlers

import (
	"net/http"

	"bookings/pkg/config"
	"bookings/pkg/models"
	"bookings/pkg/render"
)

// repository pattern - allow us to swap out components of site with little changes to the code that uses it
// repository type - for now config, later add db connection
type Repository struct {
	App *config.AppConfig
}

// repository used by handlers
var Repo *Repository

// creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// add receiver to functions to allow access to repo
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

/*
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
	//fmt.Fprintf(w, "This is the home page")
}
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
	/*
		sum := AddValues(3, 12)
		fmt.Fprintf(w, fmt.Sprintf("This is the about page and the sum is %d", sum))
		res, err := divideValues(7, 0)
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("error dividing %s", err))
			return
		}
		fmt.Fprintf(w, fmt.Sprintf("This is the about page and the div is %d", res))
	*
} */

/*
// upper case allows use of 'AddValues' outside package
func AddValues(x, y int) int {
	return x + y
}

// lower case first letter only allows 'divideValues' within package
func divideValues(x, y float32) (float32, error) {
	if y == 0.0 {
		return 0.0, errors.New("Cannot divide by zero") //fmt.Errorf("Cannot divide by zero")
	} else {
		return x / y, nil
	}
}
*/
