package main

import (
	"go-sample-webserver/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRouter(mux fiber.Router) {
	mux.Get("/", handlers.HomeHandler)
	mux.Get("/about", handlers.AboutHandler)
}
