package routes

import (
	"errors"

	"github.com/Ymit24/fiber-api/database"
	"github.com/Ymit24/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type UserDTO struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func createUserDTO(u models.User) UserDTO {
	return UserDTO{
		ID:       u.ID,
		Username: u.Username,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	userDto := createUserDTO(user)
	return c.Status(200).JSON(userDto)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)

	userDtos := []UserDTO{}

	for _, user := range users {
		userDto := createUserDTO(user)
		userDtos = append(userDtos, userDto)
	}
	return c.Status(200).JSON(userDtos)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	var user models.User

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	userDto := createUserDTO(user)

	return c.Status(200).JSON(userDto)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	var user models.User

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		Username string `json:"username"`
	}

	var updateUser UpdateUser

	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.Username = updateUser.Username

	database.Database.Db.Save(&user)

	userDto := createUserDTO(user)

	return c.Status(200).JSON(userDto)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	var user models.User

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully deleted user")

}

func DeleteAllUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Database.Db.Find(&users)

	for _, user := range users {
		if err := database.Database.Db.Delete(&user).Error; err != nil {
			return c.Status(400).JSON(err.Error())
		}
	}

	return c.Status(200).SendString("Successfully deleted all users.")
}
