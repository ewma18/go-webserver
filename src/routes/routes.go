package routes

import (
	"go-sample-webserver/src/handlers"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(mux fiber.Router) {
	mux.Get("/", handlers.HomeHandler)
	mux.Get("/about", handlers.AboutHandler)
	mux.Get("/search-availability", handlers.SearchAvailabilityHandler)
	mux.Post("/search-availability", handlers.PostSearchAvailabilityHandler)
	mux.Get("/make-reservation", handlers.MakeReservationHandler)
	mux.Get("/contact", handlers.ContactHandler)
	mux.Get("/generals-quarters", handlers.GeneralsRoomHandler)
	mux.Get("/majors-suite", handlers.MajorsRoomHandler)

	config := fiber.Static{
		Compress:      true,
		CacheDuration: 10 * time.Hour,
		MaxAge:        int(10 * 24 * time.Hour), //10 days
	}

	mux.Static("/static", "./resources/static", config)

}
