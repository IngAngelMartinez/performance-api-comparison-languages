package server

import (
	"compression/src/controllers"
	"compression/src/dependencies"

	"github.com/gofiber/fiber/v2"
)

func MapRoutes(app *fiber.App) {

	apiRouter := app.Group("/api")

	controllers := dependencies.CompressionDependencies()

	compressionRoute(apiRouter, controllers)
}

func compressionRoute(apiRouter fiber.Router, controllers *controllers.CompressionController) {

	apiRouter.Get("/compression", controllers.GetAll)

	apiRouter.Get("/compression/:id", controllers.GetById)

	apiRouter.Post("/compression/", controllers.Create)
}
