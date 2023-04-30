package core

import (
	"unknspec/core/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func InitApp() *fiber.App {
	app := fiber.New(fiber.Config{
		Views: html.New("./core/views", ".html"),
	})

	// Static files
	app.Static("/static", "./public")

	// init middleware
	// TODO: add middleware

	// init routes
	app.Get("/", routes.PostsRoute)
	app.Get("/:id<int>", routes.PostRoute)

	return app
}
