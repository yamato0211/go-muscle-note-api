package controllers

import (
	"fiber-muscles/repositories"
	"fiber-muscles/schemas"

	"github.com/gofiber/fiber/v2"
)

func CreateMenuByInputs(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token is invalid",
		})
	}
	input := new(schemas.NewMenuInput)
	if err := c.BodyParser(input); err != nil {
		return err
	}
	m, err := repositories.CreateMenu(input.Name, input.Description, input.Target, input.Weight, input.IsJoint, input.Link, userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&m)
}

func GetAllMenu(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token is invalid",
		})
	}
	m, err := repositories.GetAllMenuByID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&m)
}
