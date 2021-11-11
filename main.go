package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/alfainfdl/fiber-gorm/database"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Fiber API")
}

func main(){
	database.ConnectDb()
	app := fiber.New()
	
	app.Get("/api", welcome)

	app.Use(logger.New())
		
	log.Fatal(app.Listen(":3000"))
}