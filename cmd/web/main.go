package main

import (
	"fmt"
	"net/http"

	"github.com/Sochika/SuchGo/pkg/handlers"
)

const portNum = ":8080"

func main() {
	defer fmt.Printf("Server closed on port %s and %s is freed", portNum, portNum)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("Starting Server on port %s", portNum)
	_ = http.ListenAndServe(portNum, nil)

}
