package controllers

import (
	"fiber-muscles/repositories"
	"fiber-muscles/schemas"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateUserByInputs(c *fiber.Ctx) error {
	input := new(schemas.SignUpInput)
	if err := c.BodyParser(input); err != nil {
		return err
	}

	u, err := repositories.CreateUser(input.Name, input.Email, input.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&u)
}
