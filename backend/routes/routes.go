package routes

import (
	"AELS/middleware"
	"AELS/routes/api/auth"
	"AELS/routes/api/chat"
	"AELS/routes/api/chat/chatid"
	"AELS/routes/api/chat/chatid/complete"
	"AELS/routes/api/login"
	"AELS/routes/api/logout"
	"AELS/routes/api/user"

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

	User := Api.Group("/user", middleware.Auth())
	User.Get("/", user.GetUser)

	Auth := Api.Group("/auth", middleware.Auth())
	Auth.Get("/", auth.CheckAuth)

	Chat := Api.Group("/chat", middleware.Auth())
	Chat.Get("/", chat.ChatList)
	Chat.Post("/", chat.NewChat)
	Chat.Patch("/", middleware.ValidateQueryChatIDAndOwnership, chat.UpdateChat)
	Chat.Delete("/", middleware.ValidateQueryChatIDAndOwnership, chat.DeleteChat)

	Message := Chat.Group("/:chatid")
	Message.Use(middleware.ValidateURLChatIDAndOwnership)
	Message.Get("/", chatid.ListMessages)
	Message.Post("/", chatid.NewMessage)
	Message.Post("/complete", complete.CompleteChat())

	app.Static("/", "dist")
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./dist/index.html")
	})
}
