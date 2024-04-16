package main

import (
	"log"

	"github.com/4fort/fiber_gorm/database"
	"github.com/4fort/fiber_gorm/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func welcome(c fiber.Ctx) error {
	return c.SendString("Hello, World! OTINLKAHGDJKHAGWSDKLJJGH")
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// welcome endpoint
	api.Get("/", welcome)

	// user endpoint
	api.Post("/users", routes.CreateUser)
	api.Get("/users", routes.GetUsers)
	api.Get("/users/:id", routes.GetUser)
	api.Put("/users/:id", routes.UpdateUser)
	api.Delete("/users/:id", routes.DeleteUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
