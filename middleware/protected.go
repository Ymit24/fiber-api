package middleware

import (
	"fmt"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected() func(c *fiber.Ctx) error {
	JwtToken := []byte(os.Getenv("JWTSECRET"))
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: JwtToken, // TODO: CENTRALIZE THIS
		},
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.Status(fiber.StatusBadRequest).JSON(e.Error())
		},
	})
}

func Protected2() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("token")
		fmt.Println("Token: ", token)
		if token != "123" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}
		return c.Next()
	}
}
