package routes

import (
	"log"
	"os"
	"time"

	"github.com/Ymit24/fiber-api/database"
	"github.com/Ymit24/fiber-api/models"
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
			ID:       1,
			Username: "abc123",
		},
	})
}

func Register(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Email       string `json:"email"`
		Username    string `json:"username"`
		PasswordRaw string `json:"password_raw"`
	}

	var registerRequest RegisterRequest

	if err := c.BodyParser(&registerRequest); err != nil {
		log.Printf("\n\n[Register] FAILED: %v\n\n", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	log.Printf("\n\n[Register] request: %v\n\n", registerRequest)

	user := models.User{
		Email:       registerRequest.Email,
		PasswordEnc: registerRequest.PasswordRaw,
		Username:    registerRequest.Username,
	}

	database.Database.Db.Create(&user)

	return c.Status(fiber.StatusOK).SendString("ok")
}
