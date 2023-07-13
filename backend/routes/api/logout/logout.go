package logout

import "github.com/gofiber/fiber/v2"

// Returns 200 on success
func LogoutUser(c *fiber.Ctx) error {

	auth_cookie := new(fiber.Cookie)
	auth_cookie.Name = "auth"
	auth_cookie.Value = ""
	c.Cookie(auth_cookie)

	return c.SendStatus(200)
}
