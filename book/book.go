package book

import ("gorm.io/gorm")

func getBooks(c *fiber.Ctx){
	c.Send("All Books")
}

func getBook(c *fiber.Ctx){
	c.Send("Single Books")
}

func newBook(c *fiber.Ctx){
	c.Send("Add Books")
}

func deleteBook(c *fiber.Ctx){
	c.Send("Delete Books")
}