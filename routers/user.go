package routers

import (
	"fiber-muscles/controllers"

	"github.com/gofiber/fiber/v2"
)

func InitUserRouter(r fiber.Router) {
	r.Post("/signup", controllers.CreateUserByInputs)
}
