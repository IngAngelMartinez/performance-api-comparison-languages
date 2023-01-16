package main

import (
	"compression/src/server"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := fiber.Config{
		Immutable: true,
	}

	app := fiber.New(config)

	server.MapRoutes(app)

	app.Listen(":3000")
}
