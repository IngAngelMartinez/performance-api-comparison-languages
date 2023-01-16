package routes

import (
	"crud/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func MapRoutes(app *fiber.App) {

	api := app.Group("/api")

	productsRouter := api.Group("/products")

	mapProductsRouter(productsRouter)
}

func mapProductsRouter(productsRouter fiber.Router) {

	productsRouter.Get("/", controllers.GetAll)

	productsRouter.Get("/:id", controllers.Get)

	productsRouter.Post("/", controllers.Create)
}
