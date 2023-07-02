package routes

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var loginRequest LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse json",
		})
	}

	if loginRequest.Email != "myemail" || loginRequest.Password != "123" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // 1 week

	JwtToken := os.Getenv("JWTSECRET")

	s, err := token.SignedString([]byte(JwtToken))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"user": UserDTO{
			ID:        1,
			FirstName: "1",
			LastName:  "2",
		},
	})
}
