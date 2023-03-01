package routers

// パッケージのインポート
import (
	"github.com/gofiber/fiber/v2"
)

func InitHelloWorld(r fiber.Router) {
	r.Get("/", helloWorld)
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
