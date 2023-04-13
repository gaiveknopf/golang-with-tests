package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := NewFiberServer()

	err := app.Listen(":3336")
	if err != nil {
		return
	}
}

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func NewFiberServer() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		err := HelloWorld(c)
		if err != nil {
			return nil
		}
		return nil
	})
	return app
}
