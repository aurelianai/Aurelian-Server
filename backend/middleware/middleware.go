package middleware

import (
	"os"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Checks Tokens
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		TokenLookup:    "cookie:auth",
		SuccessHandler: SetUid,
	})
}

// Sets uint UID in Locals on token validation success
func SetUid(c *fiber.Ctx) error {
	uid := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["uid"].(float64)
	c.Locals("uid", uint(uid))
	return c.Next()
}

// Resets token expiration if valid
func Refresh() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		TokenLookup:    "cookie:auth",
		ErrorHandler:   IgnoreBadToken,
		SuccessHandler: TokenRefresh,
	})
}

func TokenRefresh(c *fiber.Ctx) error {
	c.ClearCookie("auth")

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	claims = jwt.MapClaims{
		"uid": claims["uid"],
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	auth_cookie := new(fiber.Cookie)
	auth_cookie.Name = "auth"
	auth_cookie.Value = t
	auth_cookie.Expires = time.Now().Add(time.Hour * 24)
	auth_cookie.SameSite = "Strict"
	auth_cookie.HTTPOnly = true
	c.Cookie(auth_cookie)

	return c.Next()
}

func IgnoreBadToken(c *fiber.Ctx, err error) error {
	return c.Next()
}
