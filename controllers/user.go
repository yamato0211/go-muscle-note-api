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

func LoginUserByInputs(c *fiber.Ctx) error {
	input := new(schemas.SignInInput)
	if err := c.BodyParser(input); err != nil {
		return err
	}

	t, err := repositories.UserLogin(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&t)
}

func GetMe(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token is invalid",
		})
	}

	u, err := repositories.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&u)
}
