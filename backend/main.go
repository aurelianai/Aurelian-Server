package main

import (
	"AELS/persistence"
	"AELS/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	persistence.InitAndMigrate()

	app := fiber.New()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":2140"))
}
