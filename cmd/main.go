package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Nissekissen/GO-testing/database"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	setupRoutes(app)


	app.Listen(":5000")
}
