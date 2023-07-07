package login

import (
	"AELS/persistence"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Unauthorized (401) if Uname or Pass is incorrect, 500 for all others
func LoginUser(c *fiber.Ctx) error {
	email := c.FormValue("email")
	pass := c.FormValue("pass")

	var user persistence.User
	res := persistence.DB.Where(&persistence.User{Email: &email, Password: &pass}).First(&user)
	if res.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := jwt.MapClaims{
		"uid": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	auth_cookie := new(fiber.Cookie)
	auth_cookie.Name = "auth"
	auth_cookie.Value = t
	auth_cookie.Expires = time.Now().Add(time.Hour * 24)
	auth_cookie.HTTPOnly = true
	auth_cookie.SameSite = "Strict"
	c.Cookie(auth_cookie)

	return c.SendStatus(fiber.StatusOK)
}
