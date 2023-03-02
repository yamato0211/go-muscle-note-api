package routers

import (
	"fiber-muscles/controllers"
	"fiber-muscles/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitUserRouter(r fiber.Router) {
	r.Post("/signup", controllers.CreateUserByInputs)
	r.Post("/signin", controllers.LoginUserByInputs)
	r.Use(middleware.JwtAuthMiddleware())
	r.Get("/me", controllers.GetMe)
}
