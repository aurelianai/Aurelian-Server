package main

import (
	"AELS/api/chat"
	"AELS/api/login"
	"AELS/api/logout"
	"AELS/middleware"
	"AELS/persistence"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	persistence.InitAndMigrate()

	app := fiber.New()

	SetupRoutes(app)

	log.Fatal(app.Listen(":2140"))
}

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	app.Static("/", "/dist")

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Aurelian Enterprise Language Server ready to recieve requests.")
	})

	api := app.Group("/api")
	api.Post("/login", login.LoginUser)
	api.Post("/logout", logout.LogoutUser)

	chat_group := api.Group("/chat")
	chat_group.Use(middleware.Protected())
	chat_group.Get("/", chat.ChatList)
	chat_group.Post("/", chat.NewChat)
	chat_group.Patch("/", chat.UpdateChat)
	chat_group.Delete("/", chat.DeleteChat)
	chat_group.Use(middleware.Refresh())
}
