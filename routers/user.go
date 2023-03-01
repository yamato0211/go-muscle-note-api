package routers

import (
	"fiber-muscles/controllers"
	"fiber-muscles/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitUserRouter(r fiber.Router) {
	r.Post("/signup", middleware.TestMiddleware(), controllers.CreateUserByInputs)
}
