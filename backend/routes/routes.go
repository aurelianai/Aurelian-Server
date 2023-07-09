package routes

import (
	"AELS/middleware"
	"AELS/routes/api/auth"
	"AELS/routes/api/chat"
	"AELS/routes/api/chat/id/message"
	"AELS/routes/api/login"
	"AELS/routes/api/logout"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).SendString(
			"Aurelian Enterprise Language Server ready to recieve requests.",
		)
	})

	Api := app.Group("/api")
	Api.Post("/login", login.LoginUser)
	Api.Post("/logout", logout.LogoutUser)

	Auth := Api.Group("/auth", middleware.Auth())
	Auth.Get("/", auth.CheckAuth)

	Chat := Api.Group("/chat", middleware.Auth())
	Chat.Get("/", chat.ChatList)
	Chat.Post("/", chat.NewChat)
	Chat.Patch("/", middleware.ValidateQueryChatIDAndOwnership, chat.UpdateChat)
	Chat.Delete("/", middleware.ValidateQueryChatIDAndOwnership, chat.DeleteChat)

	Message := Chat.Group("/:chatid")
	Message.Use(middleware.ValidateURLChatIDAndOwnership)
	Message.Get("/", message.ListMessages)
	Message.Post("/", message.NewMessage)

	app.Static("/", "dist")
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./dist/index.html")
	})
}
