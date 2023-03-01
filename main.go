package main

import (
	"fiber-muscles/config"
	"fiber-muscles/models"
	"fiber-muscles/routers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.LoadConfig()
	models.InitDB()
	routers.Router(app)

	app.Listen(fmt.Sprintf(":%s", config.ApiPort))
}
