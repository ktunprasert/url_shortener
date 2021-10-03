package main

import (
	"github.com/gofiber/fiber"
	"github.com/ktunprasert/url_shortener/short_db"
)

func main() {
	app := fiber.New()

	app.Get("/:short", func(c *fiber.Ctx) error {
		short := c.Params("short")

		return c.SendString("Hello, World 👋! " + short)
	})

	short_db.Setup()

	app.Listen(":3000")
}
