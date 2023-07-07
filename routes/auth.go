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

	var user models.User
	database.Database.Db.Find(&user, "email = ?", loginRequest.Email)

	// TODO: HANDLE SALTING
	if user.ID == 0 || user.PasswordEnc != loginRequest.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
	}

	s, err := makeJwt()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate jwt",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"user": UserDTO{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
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

func makeJwt() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // 1 week

	JwtToken := os.Getenv("JWTSECRET")

	s, err := token.SignedString([]byte(JwtToken))
	if err != nil {
		return "", err
	}
	return s, nil
}
