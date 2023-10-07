package main

import (
	"github.com/Nissekissen/GO-testing/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Post("/facts", handlers.CreateFact)

	app.Get("/login", handlers.Login)
	app.Get("/callback", handlers.Callback)
}
