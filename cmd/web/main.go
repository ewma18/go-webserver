package main

import (
	"fmt"
	"go-sample-webserver/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	fmt.Printf("Starting application on port %s \n", portNumber)
	fmt.Println("Listening...")

	http.ListenAndServe(portNumber, nil)
}
