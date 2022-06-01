package main

import (
	"basic-web/pkg/config"
	"basic-web/pkg/handlers"
	"basic-web/pkg/renders"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig
	//# creates a new template cache
	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	//# stores the template cache in the app config
	app.TemplateCache = tc
	app.UseCache = false //@ cache everytime something changes

	//# creates new repository and sets it in the app config
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//# creates the new template cache
	renders.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
