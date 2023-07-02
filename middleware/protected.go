package middleware

import (
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
