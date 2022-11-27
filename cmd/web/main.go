package main

import (
	"fmt"
	"log"
	"github.com/Ramnath-Karthikesan/bookings/pkg/config"
	"github.com/Ramnath-Karthikesan/bookings/pkg/handlers"
	"github.com/Ramnath-Karthikesan/bookings/pkg/render"
	"net/http"
	"time"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	// fmt.Println("Hello world!")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello World!")
	// 	fmt.Println("Number of bytes written", n)

	// 	if err != nil{
	// 		fmt.Println(err)
	// 	}
	// })


	// change this to true when in production
	app.InProduction = false
	
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/divide", handlers.Divide)

	fmt.Println("Starting application on port", portNumber)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

	// fmt.Println("Starting application on port", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)
}
