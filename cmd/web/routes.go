package main

import (
	"go-sample-webserver/pkg/config"
	"go-sample-webserver/pkg/handlers"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func setupRouter(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.RealIP)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Use(middleware.Timeout(3 * time.Second))
	mux.Use(middleware.Compress(5))

	mux.Get("/", handlers.HomeHandler)
	mux.Get("/about", handlers.AboutHandler)

	return mux
}
