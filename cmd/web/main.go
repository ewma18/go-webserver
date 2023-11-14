package main

import (
	"fmt"
	"go-sample-webserver/pkg/config"
	"go-sample-webserver/pkg/handlers"
	"go-sample-webserver/pkg/renders"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	setup()

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	fmt.Printf("Starting application on port %s \n", portNumber)
	fmt.Println("Listening...")

	http.ListenAndServe(portNumber, nil)
}

func setup() {
	config := config.GetAppConfig()
	config.UseCache = true

	if config.UseCache {
		templateCache, err := renders.PreLoadTemplates()
		if err != nil {
			log.Fatal("cannot create template cache", err)
		}
		config.TemplateCache = templateCache
	}
}
