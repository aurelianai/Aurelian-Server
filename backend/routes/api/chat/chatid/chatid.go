package chatid

import (
	"AELS/persistence"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Returns all messages in chat
func ListMessages(c *fiber.Ctx) error {
	chatid := c.Locals("chatid").(uint64)

	var messages []persistence.Message
	if err := persistence.DB.Where("chat_id = ?", chatid).Find(&messages).Error; err != nil {
		err := fmt.Sprintf("Error Collecting Messages: %s", err.Error())
		fmt.Println(err)
		return c.Status(500).SendString(err)
	}

	return c.JSON(messages)
}

// Creates a new message in chat
func NewMessage(c *fiber.Ctx) error {
	chatid := c.Locals("chatid").(uint64)

	var new_message = new(persistence.Message)
	if err := c.BodyParser(&new_message); err != nil {
		err := fmt.Sprintf("Error Parsing message body: %s", err.Error())
		fmt.Println(err)
		return c.Status(400).SendString(err)
	}
	new_message.ChatID = chatid

	if err := persistence.DB.Create(&new_message).Error; err != nil {
		err := fmt.Sprintf("Error creating new message: %s", err.Error())
		fmt.Println(err)
		return c.Status(500).SendString(err)
	}

	return c.JSON(new_message)
}
