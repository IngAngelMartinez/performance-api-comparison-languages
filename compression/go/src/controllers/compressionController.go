package controllers

import (
	"compression/src/models"
	"compression/src/services"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type CompressionController struct {
	CompressionService *services.CompressionService
}

func (ctr *CompressionController) GetAll(c *fiber.Ctx) error {

	result, err := ctr.CompressionService.GetAll()

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if result == nil {

		return c.Status(fiber.StatusNotFound).JSON([]models.Compression{})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (ctr *CompressionController) GetById(c *fiber.Ctx) error {

	id := c.Params("id")

	compression, err := ctr.CompressionService.GetById(id)

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if reflect.DeepEqual(compression, models.Compression{}) {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
	}

	return c.Status(fiber.StatusOK).JSON(compression)
}

func (ctr *CompressionController) Create(c *fiber.Ctx) error {

	var compression models.Compression

	if err := c.BodyParser(&compression); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	id, err := ctr.CompressionService.Create(compression)

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}
