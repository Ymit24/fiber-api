package routes

import (
	"errors"
	"time"

	"github.com/Ymit24/fiber-api/database"
	"github.com/Ymit24/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

/*

{
    id: 1,
    user: {
        id: 22,
        first_name: "a",
        last_name: "b"
    },
    product: {
        id: 1,
        name: "object",
        serial_number: "12345"
    }
}


*/

type OrderDTO struct {
	ID        uint       `json:"id"`
	User      UserDTO    `json:"user"`
	Product   ProductDTO `json:"product"`
	CreatedAt time.Time  `json:"order_date"`
}

func createOrderDto(order models.Order, userDto UserDTO, productDto ProductDTO) OrderDTO {
	return OrderDTO{
		ID:        order.ID,
		User:      userDto,
		Product:   productDto,
		CreatedAt: order.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserId, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductId, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	userDto := createUserDTO(user)
	productDto := createProductDTO(product)
	orderDto := createOrderDto(order, userDto, productDto)

	return c.Status(200).JSON(orderDto)
}

func findOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

func DeleteOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var order models.Order
	if err := findOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&order).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully deleted order")

}
