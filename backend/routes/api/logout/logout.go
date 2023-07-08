package logout

import "github.com/gofiber/fiber/v2"

// Returns 200 on success
func LogoutUser(c *fiber.Ctx) error {
	c.ClearCookie()
	return c.SendStatus(200)
}
