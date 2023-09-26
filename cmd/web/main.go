package main

import (
	"fmt"
	"github.com/rozdolsky33/Booking-App/config"
	"github.com/rozdolsky33/Booking-App/pkg/handlers"
	"github.com/rozdolsky33/Booking-App/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	// get teh template cache from the app config
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UserCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
