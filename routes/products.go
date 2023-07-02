package routes

import (
	"errors"

	"github.com/Ymit24/fiber-api/database"
	"github.com/Ymit24/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type ProductDTO struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func createProductDTO(p models.Product) ProductDTO {
	return ProductDTO{
		ID:           p.ID,
		Name:         p.Name,
		SerialNumber: p.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	productDto := createProductDTO(product)

	return c.Status(200).JSON(productDto)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)

	productDtos := []ProductDTO{}
	for _, product := range products {
		productDto := createProductDTO(product)
		productDtos = append(productDtos, productDto)
	}

	return c.Status(200).JSON(productDtos)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("product does not exist")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	productDto := createProductDTO(product)
	return c.Status(200).JSON(productDto)

}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	var product models.Product

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateProduct UpdateProduct

	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	product.Name = updateProduct.Name
	product.SerialNumber = updateProduct.SerialNumber

	database.Database.Db.Save(&product)

	productDto := createProductDTO(product)

	return c.Status(200).JSON(productDto)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	var product models.Product

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully deleted product")

}
