package user

import (
	"AELS/persistence"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	userId := c.Locals("uid").(uint64)

	var u persistence.User
	if err := persistence.DB.Where("id = ?", userId).First(&u).Error; err != nil {
		return err
	}

	return c.JSON(u)
}
