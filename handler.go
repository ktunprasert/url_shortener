package main

import (
	"github.com/gofiber/fiber"
	"github.com/ktunprasert/url_shortener/short_db"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/list", func(c *fiber.Ctx) error {
		return c.JSON(short_db.ListShort())
	})

	app.Get("/:short", func(c *fiber.Ctx) error {
		url := short_db.ReadShort(c.Params("short"))
		if url != "" {
			return c.Redirect(url, 301)
		}
		return c.Status(404).SendString("Not found!")
	})

	app.Post("/:short", func(c *fiber.Ctx) error {
		return c.JSON(short_db.WriteShort(c.Params("short")))
	})

}

func main() {
	app := fiber.New()

	SetupRoutes(app)

	short_db.SetupDB()

	app.Listen(":3000")
}
