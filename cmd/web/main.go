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
	var app config.SystemConfig

	pages, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = pages

	render.RefreshTemplates(&app)

	defer fmt.Printf("Server closed on port %s and %s is freed", portNum, portNum)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("Starting Server on port %s", portNum)
	_ = http.ListenAndServe(portNum, nil)

}
