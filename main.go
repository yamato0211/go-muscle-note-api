package main

import (
	"fiber-muscles/config"
	"fiber-muscles/routers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	conf := config.LoadConfig()
	routers.Router(app)

	app.Listen(fmt.Sprintf(":%s", conf.Server.Port))
}
