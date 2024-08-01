package main

import (
	"fmt"
	"go-sample-webserver/src/config"
	"go-sample-webserver/src/renders"
	"go-sample-webserver/src/routes"
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
	config.InProduction = false

	server := fiber.New(fiber.Config{
		AppName:      "Go Sample Webserver (Go Fiber)",
		ServerHeader: "Fiber",
		Immutable:    true,
	})

	// Use global middlewares.
	server.Use(requestid.New())
	server.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:       "[${time}] ${ip} ${locals:requestid} - ${method} ${path} -${latency} ${status} ${red}${error}​${reset}\n",
		TimeFormat:   "2006-01-01 15:04:05.000",
		TimeInterval: 500 * time.Millisecond,
	}))

	//app.Use(logger.New())

	server.Use(cors.New())
	server.Use(compress.New())
	server.Use(etag.New())

	server.Use(favicon.New())
	server.Use(recover.New())
	//app.Use(middlewares.NoSurf)

	renders.SetupPageTemplates(config)
	routes.SetupRouter(server.Group("/"))

	fmt.Printf("Starting application on port %s \n", portNumber)
	fmt.Println("Listening...")

	server.Listen(portNumber)

	//http.ListenAndServe(portNumber, nil)
}
