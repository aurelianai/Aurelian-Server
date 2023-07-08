package message

import (
	"AELS/persistence"
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Validates ChatID Parameter
// Verifies Chat Exists, User Owns it
func ListMessages(c *fiber.Ctx) error {
	UID := c.Locals("uid").(uint)
	ChatID, err := strconv.ParseUint(c.Params("ChatID"), 10, 64)
	if err != nil {
		err := fmt.Sprintf("Recieved invalid ChatID param: %s\n", c.Params("ChatID"))
		fmt.Println(err)
		return c.Status(400).SendString(err)
	}

	var chat persistence.Chat
	err = persistence.DB.Where("chat_id = ?", ChatID).First(&chat).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("Chat id(%d) requested that doesn't exist\n", ChatID)
		return c.Status(404).SendString(fmt.Sprintf("Chat id(%d) doesn't exist", ChatID))
	} else if err != nil {
		fmt.Println(err.Error())
		return c.Status(500).SendString(err.Error())
	}

	if chat.UserID != c.Locals("uid").(uint) {
		err := fmt.Sprintf("User with uid(%d) attempted to access chat id(%d) they do not own", ChatID, UID)
		fmt.Println(err)
		return c.Status(401).SendString(err)
	}

	var messages []persistence.Message
	if err = persistence.DB.Where("chat_id = ?", ChatID).Find(&messages).Error; err != nil {
		err := fmt.Sprintf("Error Collecting Messages: %s\n", err.Error())
		fmt.Println(err)
		return c.Status(500).SendString(err)
	}

	return c.JSON(messages)
}
