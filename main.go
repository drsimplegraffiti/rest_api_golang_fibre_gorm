package main

import (
	"log"

	"github.com/drsimplegraffiti/fibre-api/database"
	"github.com/drsimplegraffiti/fibre-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to golang rest api")
}

func setupRoute(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", welcome)

	//User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUser)
	app.Get("/api/users/:id", routes.GetSingleUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	//Product endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetSingleProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)

	//Order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetSingleOrder)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoute(app)

	log.Fatal(app.Listen(":3002"))
}
