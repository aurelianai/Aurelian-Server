package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Service Healthy")
	})

	app.Static("/", "/dist")

	log.Fatal(app.Listen(":2140"))
}
