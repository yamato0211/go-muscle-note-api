package routers

// パッケージのインポート
import (
	"github.com/gofiber/fiber/v2"
)

func Router(e *fiber.App) {
	api := e.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			InitHelloWorld(v1)
		}
	}
}
