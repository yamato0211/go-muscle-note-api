package middleware

import (
	"fiber-muscles/services"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
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

func JwtAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		h := c.Get("Authorization")
		if h != "" {
			ha := strings.Split(h, " ")
			if len(ha) == 2 {
				if ha[0] == "Bearer" {
					t, err := services.ValidateJwt(ha[1])
					if claims, ok := t.Claims.(*jwt.MapClaims); ok && t.Valid {
						userId := (*claims)["sub"].(string)
						c.Locals("user_id", userId)
					} else {
						fmt.Println("test1")
						fmt.Println(err)
						return err
					}
				}
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON((fiber.Map{
					"error": "Bearer stirng not found",
				}))
			}
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization Header not found",
			})
		}

		return c.Next()
	}
}
