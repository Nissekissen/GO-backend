package main

import (
	"github.com/Nissekissen/GO-testing/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":5000")
}
