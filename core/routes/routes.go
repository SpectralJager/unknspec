package routes

import (
	"strconv"
	"unknspec/core/db"

	"github.com/gofiber/fiber/v2"
)

func PostsRoute(ctx *fiber.Ctx) error {
	posts, _ := db.GetPosts()
	return ctx.Render("posts", fiber.Map{
		"Posts": posts,
	})
}

func PostRoute(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	res, _ := strconv.ParseInt(id, 10, 32)
	post, err := db.GetPostById(int(res))
	if err != nil {
		return ctx.SendString(err.Error())
	}
	return ctx.Render("post", fiber.Map{
		"Post": post,
	})
}
