package main

import (
	"crud/src/database"
	"crud/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Mongodb()

	app := fiber.New()

	routes.MapRoutes(app)

	app.Listen(":3000")
}
