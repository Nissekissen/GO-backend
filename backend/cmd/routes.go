package main

import (
	"github.com/Nissekissen/GO-testing/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Get("/user", handlers.Authenticate, handlers.GetUser)

	app.Get("/login", handlers.Login)
	app.Get("/callback", handlers.Callback)
	app.Get("/refresh", handlers.Refresh)

	app.Get("/posts", handlers.GetPosts)
	app.Get("/posts/:id", handlers.GetPost)
	app.Post("/posts", handlers.Authenticate, handlers.CreatePost)
	app.Delete("/posts/:id", handlers.Authenticate, handlers.DeletePost)

}
