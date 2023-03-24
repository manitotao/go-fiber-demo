package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manitotao/go-fiber-demo/book"
)

func setupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/book", book.GetBooks)
	v1.Get("/book/:id", book.GetBook)
	v1.Post("/book", book.NewBook)
	v1.Delete("/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	setupRoutes(app)
	app.Listen(":3000")
}
