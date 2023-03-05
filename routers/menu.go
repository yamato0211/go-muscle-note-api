package routers

import (
	"fiber-muscles/controllers"
	"fiber-muscles/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitMenuRouter(r fiber.Router) {
	r.Use(middleware.JwtAuthMiddleware())
	r.Get("/", controllers.GetAllMenu)
	r.Post("/", controllers.CreateMenuByInputs)
}
