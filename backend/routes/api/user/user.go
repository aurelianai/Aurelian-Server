package user

import (
	"AELS/persistence"
	"fmt"
	"os"

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

/*
# Schema

{ email: string } --> { password: string } | error.

# Notes

Automatically generates a 16 character secure password, adds user into the database
*/
func CreateUser(c *fiber.Ctx) error {

	if c.Get("Authorization") != os.Getenv("USER_CREATE_ACCESS_KEY") {
		fmt.Printf("Got %s, but needed %s", c.Get("Authorization"), os.Getenv("USER_CREATE_ACCESS_KEY"))
		return c.SendStatus(403)
	}

	type UserSignUpPayload struct {
		Email string
	}

	var userSignUpPayload UserSignUpPayload
	if err := c.BodyParser(&userSignUpPayload); err != nil {
		return err
	}

	u := new(persistence.User)
	u.Email = userSignUpPayload.Email
	u.Password = "didntcreatepasswordyet!"
	if err := persistence.DB.Create(u).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{"password": u.Password})
}
