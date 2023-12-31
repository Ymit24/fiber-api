package main

import (
	"log"

	"github.com/Ymit24/fiber-api/database"
	"github.com/Ymit24/fiber-api/middleware"
	"github.com/Ymit24/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the API")
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", welcome)

	users := api.Group("/users")
	users.Post("/", routes.CreateUser)
	users.Get("/", routes.GetUsers)
	users.Get("/secret", middleware.Protected(), routes.GetUsers)
	users.Get("/:id", routes.GetUser)
	users.Put("/:id", routes.UpdateUser)
	users.Delete("/:id", routes.DeleteUser)
	users.Delete("/debug/delete-all", routes.DeleteAllUsers)

	products := api.Group("/products")
	products.Post("/", routes.CreateProduct)
	products.Get("/", routes.GetProducts)
	products.Get("/:id", routes.GetProduct)
	products.Put("/:id", routes.UpdateProduct)
	products.Delete("/:id", routes.DeleteProduct)

	orders := api.Group("/orders")
	orders.Post("/", routes.CreateOrder)
	orders.Delete("/:id", routes.DeleteOrder)

	auth := api.Group("/auth")
	auth.Post("/login", routes.Login)
	auth.Post("/register", routes.Register)
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file.")
	}

	database.ConnectDb()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
