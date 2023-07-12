package login

import (
	"AELS/persistence"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type LoginPayload struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

// Unauthorized (401) if Uname or Pass is incorrect, 500 for all others
func LoginUser(c *fiber.Ctx) error {

	var login_payload LoginPayload
	if err := c.BodyParser(&login_payload); err != nil {
		fmt.Println(err.Error())
		return c.Status(500).SendString(err.Error())
	}

	var user persistence.User
	err := persistence.DB.Where("email = ? AND password = ?", &login_payload.Email, &login_payload.Pass).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.SendStatus(404)
	} else if err != nil {
		fmt.Println(err.Error())
		return c.Status(500).SendString(err.Error())
	}

	claims := jwt.MapClaims{
		"uid": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err.Error())
		c.SendStatus(500)
	}

	auth_cookie := new(fiber.Cookie)
	auth_cookie.Name = "auth"
	auth_cookie.Value = t
	auth_cookie.Expires = time.Now().Add(time.Hour * 24)
	auth_cookie.HTTPOnly = true
	auth_cookie.SameSite = "Strict"
	c.Cookie(auth_cookie)

	return c.SendStatus(200)
}
