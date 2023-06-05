package main

import (
	"context"
	"log"
	"time"
	"unknspec/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var db = database.NewSQLite("./src/database/db.sqlite")

func main() {
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		AppName:      "Unknspec website",
		ServerHeader: "Unknspec",
		Concurrency:  10,
		GETOnly:      true,
		Views:        engine,
	})

	app.Static("/public", "./public")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/posts", fiber.StatusPermanentRedirect)
	})

	app.Get("/posts", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		tags, err := db.TagAll(ctx)
		if err != nil {
			return err
		}
		articles, err := db.ArticleAll(ctx, 0, 0)
		if err != nil {
			return err
		}
		for _, article := range articles {
			article.Tags, err = db.GetTagsOfArticleId(ctx, article.Id)
			if err != nil {
				return err
			}
		}
		return c.Render("posts", &fiber.Map{
			"tags":     tags,
			"articles": articles,
		})
	})

	app.Get("/posts/:id", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		id, err := c.ParamsInt("id", 0)
		if err != nil {
			return err
		}
		article, err := db.ArticleGetById(ctx, id)
		if err != nil {
			return err
		}
		return c.Render("post", &fiber.Map{
			"article": article,
		})
	})

	app.Get("/tag/:name", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		name := c.Params("name")
		if name == "" {
			return c.Redirect("/", fiber.StatusTemporaryRedirect)
		}
		tag, err := db.TagGetByName(ctx, name)
		if err != nil {
			return err
		}
		articles, err := db.GetArticlesFromTagsId(ctx, tag.Id)
		if err != nil {
			return err
		}
		for _, article := range articles {
			article.Tags, err = db.GetTagsOfArticleId(ctx, article.Id)
			if err != nil {
				return err
			}
		}
		return c.Render("tag", &fiber.Map{
			"tag":      tag,
			"articles": articles,
		})
	})

	log.Println("Starting server...")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
