package routes

import "github.com/gofiber/fiber/v2"

// '/'
func PostsRoute(ctx *fiber.Ctx) error {
	return ctx.Render("base", fiber.Map{
		"Msg": "Posts",
	})
}

// '/apps'
func AppsRoute(ctx *fiber.Ctx) error {
	return ctx.SendString("Apps")
}
