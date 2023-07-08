package auth

import "github.com/gofiber/fiber/v2"

func CheckAuth(c *fiber.Ctx) error {
	return c.Status(200).SendString("JWT Valid")
}
