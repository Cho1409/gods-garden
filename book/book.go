package book

import "github.com/gofiber/fiber"

// You get all books
func GetBooks(c *fiber.Ctx) {
	c.Send("All Books")
}

// You get a single book
func GetBook(c *fiber.Ctx) {
	c.Send("A Single Book")
}

// Add a New Book to the Database
func NewBook(c *fiber.Ctx) {
	c.Send("Adds a new Book")
}

// Delete book from Database
func DeleteBook(c *fiber.Ctx) {
	c.Send("Deletes a Book")
}
