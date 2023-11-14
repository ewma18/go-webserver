package main

import (
	"fmt"
	"go-sample-webserver/pkg/config"
	"go-sample-webserver/pkg/renders"
	"net/http"
)

const portNumber = ":8080"

func main() {
	config := config.GetAppConfig()
	//config.UseCache = true

	renders.SetupPageTemplates(config)
	router := setupRouter(config)

	server := &http.Server{
		Addr:    portNumber,
		Handler: router,
	}

	fmt.Printf("Starting application on port %s \n", portNumber)
	fmt.Println("Listening...")

	server.ListenAndServe()

	//http.ListenAndServe(portNumber, nil)
}
