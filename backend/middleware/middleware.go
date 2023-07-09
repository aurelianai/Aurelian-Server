package middleware

import (
	"AELS/persistence"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// Checks Token validitity, checks User has not been deleted, on success sets uid in c.Locals as a uint64, and refreshes token
func Auth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		TokenLookup:    "cookie:auth",
		SuccessHandler: ValidateUserAndAuthRefresh,
	})
}

// If user exists sets uid in c.Locals and refreshes auth. Only runs if token is valid
func ValidateUserAndAuthRefresh(c *fiber.Ctx) error {
	uid := uint64(c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["uid"].(float64))

	var user persistence.User
	if err := persistence.DB.First(&user, uid).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err := fmt.Sprintf("User id(%d) attempted access that does not exist", uid)
		fmt.Println(err)
		return c.Status(401).SendString(err)
	}

	c.Locals("uid", uid)

	c.ClearCookie("auth")
	claims := jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		err := fmt.Sprintf("Error signing refreshed key: %s", err.Error())
		fmt.Println(err)
		return c.Status(500).SendString(err)
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

// TODO DRY this up. Worried about passing ctx to helper function

// Validates chatid url param, ensures chat exists and that the uid from the JWT owns that chat, on success sets uint64 chatid in c.Locals
func ValidateURLChatIDAndOwnership(c *fiber.Ctx) error {
	uid := c.Locals("uid")
	chatid, err := strconv.ParseUint(c.Params("chatid"), 10, 64)
	if err != nil {
		err := fmt.Sprintf("Recieved invalid chatid param: '%s' err: %s", c.Params("chatid"), err.Error())
		fmt.Println(err)
		return c.Status(400).SendString(err)
	}

	var chat persistence.Chat
	err = persistence.DB.First(&chat, chatid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err := fmt.Sprintf("Chat with id(%d) does not exist", chatid)
		fmt.Println(err)
		return c.Status(404).SendString(err)
	} else if err != nil {
		err := fmt.Sprintf("Error retrieving chat id(%d) '%s'", chatid, err.Error())
		fmt.Println(err)
		return c.Status(500).SendString(err)
	}

	if chat.UserID != c.Locals("uid").(uint64) {
		err := fmt.Sprintf("User id(%d) attempted to access chat id(%d) which they do not own", uid, chatid)
		fmt.Println(err)
		return c.Status(401).SendString(err)
	}

	c.Locals("chatid", uint64(chatid))

	return c.Next()
}

func ValidateQueryChatIDAndOwnership(c *fiber.Ctx) error {
	uid := c.Locals("uid")
	chatid, err := strconv.ParseUint(c.Query("chatid"), 10, 64)
	if err != nil {
		err := fmt.Sprintf("Recieved invalid ChatID param: %s err: %s", c.Params("ChatID"), err.Error())
		fmt.Println(err)
		return c.Status(400).SendString(err)
	}

	var chat persistence.Chat
	err = persistence.DB.First(&chat, chatid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err := fmt.Sprintf("Chat with id(%d) does not exist", chatid)
		fmt.Println(err)
		return c.Status(404).SendString(err)
	} else if err != nil {
		err := fmt.Sprintf("Error retrieving chat id(%d) '%s'", chatid, err.Error())
		fmt.Println(err)
		return c.Status(500).SendString(err)
	}

	if chat.UserID != c.Locals("uid").(uint64) {
		err := fmt.Sprintf("User id(%d) attempted to access chat id(%d) which they do not own", uid, chatid)
		fmt.Println(err)
		return c.Status(401).SendString(err)
	}

	c.Locals("chatid", uint64(chatid))

	return c.Next()
}
