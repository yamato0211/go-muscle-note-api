package controllers

import (
	"fiber-muscles/repositories"
	"fiber-muscles/schemas"

	"github.com/gofiber/fiber/v2"
)

func CreateUserByInputs(c *fiber.Ctx) error {
	input := new(schemas.SignUpInput)
	if err := c.BodyParser(input); err != nil {
		return err
	}

	u, err := repositories.CreateUser(input.Name, input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(&u)
}
