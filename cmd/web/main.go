package main

import (
	"fmt"
	"go-sample-webserver/pkg/config"
	"go-sample-webserver/pkg/renders"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

const portNumber = ":8080"

func main() {
	config := config.GetAppConfig()
	//config.UseCache = true

	app := fiber.New(fiber.Config{
		AppName:      "Go Sample Webserver (Go Fiber)",
		ServerHeader: "Fiber",
		Immutable:    true,
	})

	// Use global middlewares.
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:       "[${time}] ${ip} ${locals:requestid} - ${method} ${path} -${latency} ${status} ${red}${error}​${reset}\n",
		TimeFormat:   "2006-01-01 15:04:05.000",
		TimeInterval: 500 * time.Millisecond,
	}))

	//app.Use(logger.New())

	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(recover.New())

	renders.SetupPageTemplates(config)
	setupRouter(app.Group("/"))

	fmt.Printf("Starting application on port %s \n", portNumber)
	fmt.Println("Listening...")

	app.Listen(portNumber)

	//http.ListenAndServe(portNumber, nil)
}
