package routes

import (
	"AELS/middleware"
	"AELS/routes/api/auth"
	"AELS/routes/api/chat"
	"AELS/routes/api/login"
	"AELS/routes/api/logout"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	app.Static("/", "/dist")

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString(
			"Aurelian Enterprise Language Server ready to recieve requests.",
		)
	})

	Api := app.Group("/api")
	Api.Post("/login", login.LoginUser)
	Api.Post("/logout", logout.LogoutUser)

	Auth := Api.Group("/auth", middleware.Protected())
	Auth.Get("/", auth.CheckAuth)

	Chat := Api.Group("/chat", middleware.Protected())
	Chat.Get("/", chat.ChatList)
	Chat.Post("/", chat.NewChat)
	Chat.Patch("/", chat.UpdateChat)
	Chat.Delete("/", chat.DeleteChat)
	Chat.Use(middleware.Refresh())
}
