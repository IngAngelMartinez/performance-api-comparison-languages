package controllers

import (
	"crud/src/models"
	"crud/src/services"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {

	products, err := services.GetAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return c.Status(fiber.StatusOK).JSON(products)
}

func Get(c *fiber.Ctx) error {

	id := c.Params("id")

	product, err := services.Get(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	if product == (models.Product{}) {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func Create(c *fiber.Ctx) error {

	var product models.Product

	err := c.BodyParser(&product)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid body"})
	}

	id, err := services.Create(product)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}
