package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sochika/SuchGo/pkg/config"
	"github.com/Sochika/SuchGo/pkg/handlers"
	"github.com/Sochika/SuchGo/pkg/render"
)

const portNum = ":8080"

func main() {
	defer fmt.Printf("Server closed on port %s, %s is freed", portNum, portNum)
	var app config.SystemConfig

	pages, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("Cannot create template cache: %v", err)
	}

	app.TemplateCache = pages
	app.UseCache = false

	render.RefreshTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("Starting Server on port %s \n", portNum)
	_ = http.ListenAndServe(portNum, nil)

}
