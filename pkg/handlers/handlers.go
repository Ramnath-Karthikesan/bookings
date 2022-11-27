package handlers

import (
	"errors"
	"fmt"
	"github.com/Ramnath-Karthikesan/bookings/pkg/config"
	"github.com/Ramnath-Karthikesan/bookings/pkg/models"
	"github.com/Ramnath-Karthikesan/bookings/pkg/render"
	"net/http"
)

// TemplateData  holds data sent from handlers to templates

// Repo the repository used by the handlers
var Repo *Repository

// Repository  is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the home page")
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "index.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)

	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP :=  m.App.Session.GetString(r.Context(), "remote_ip")

	
	stringMap["remote_ip"] = remoteIP
	//send data to template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Divide is the divide page handler
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 10.0, f)
}

//addValues add two integers and return the sum
// func addValues(x, y int) int {
// 	return x + y
// }

// divideValues divides 2 numbers
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
