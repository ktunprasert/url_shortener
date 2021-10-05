package main

import (
	"github.com/gofiber/fiber"
	"github.com/ktunprasert/url_shortener/short_db"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/list_short", func(c *fiber.Ctx) error {
		return c.JSON(short_db.ListShort())
	})

	app.Get("/list_stat", func(c *fiber.Ctx) error {
		return c.JSON(short_db.ListStat())
	})

	app.Get("/:short", func(c *fiber.Ctx) error {
		short := c.Params("short")
		url := short_db.ReadShort(short)
		if url != "" {
			_ = short_db.WriteStat(short, c.IP())
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
