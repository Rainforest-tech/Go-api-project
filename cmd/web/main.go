package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rainforest-tech/Go-api-project/pkg/config"
	"github.com/Rainforest-tech/Go-api-project/pkg/handlers"
	"github.com/Rainforest-tech/Go-api-project/pkg/render"
)

const portNumber = ":8080"

// Home is the about page handler

// addValues add 2 integer and sum it

//main is the main app func
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCash()
	if err != nil {
		log.Fatal("Cannot create template cash")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting app on port %s\n", portNumber))

	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
