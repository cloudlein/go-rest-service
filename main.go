package main

import (
	"github.com/cloudlein/go-rest-service/database"
	"github.com/cloudlein/go-rest-service/database/migrations"
	"github.com/cloudlein/go-rest-service/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func run() error {
	app := fiber.New()

	database.DatabaseInit()

	migrations.Migrations()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	routes.RoutesInit(app)

	return app.Listen(":3000")
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Warn("No .env file found")
	}

	if err := run(); err != nil {
		log.Fatal("cannot start server", err)
	}

}
