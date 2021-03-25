package main

import (
	"gods_garden/book"

	"github.com/gofiber/fiber"
)

func heyThere(c *fiber.Ctx) {
	c.Send("Hey There!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(3000)
}
