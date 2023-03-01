package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func TestMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("ルートハンドラ実行前")
		c.Next()
		fmt.Println("ルートハンドラ実行後")
		return nil
	}
}
