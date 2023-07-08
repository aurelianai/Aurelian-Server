package chat

import (
	"AELS/persistence"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ChatList(c *fiber.Ctx) error {
	var chats []persistence.Chat
	res := persistence.DB.Find(&chats)
	if res.Error != nil {
		return c.SendStatus(500)
	}
	return c.JSON(chats)
}

// If user doesn't exist, it will violate foreign key constraint
// on the UserID column
func NewChat(c *fiber.Ctx) error {
	new_chat := new(persistence.Chat)
	if err := c.BodyParser(new_chat); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	new_chat.UserID = c.Locals("uid").(uint)
	if err := persistence.DB.Create(&new_chat).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(new_chat)
}

// Checks that Chat Exists and that User owns it
func UpdateChat(c *fiber.Ctx) error {
	new_chat := new(persistence.Chat)
	if err := c.BodyParser(new_chat); err != nil {
		log.Printf("%s", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	cur_chat := new(persistence.Chat)
	if err := persistence.DB.First(&cur_chat, new_chat.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return c.SendStatus(404)
	} else if err != nil {
		log.Printf("%s", err.Error())
		return c.Status(500).SendString(err.Error())
	}
	if cur_chat.ID != c.Locals("uid").(uint) {
		return c.SendStatus(401)
	}

	if err := persistence.DB.Save(&new_chat).Error; err != nil {
		log.Printf("%s", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(&new_chat)
}

// Checks that Chat exists and that User owns it
func DeleteChat(c *fiber.Ctx) error {
	chat_to_delete := new(persistence.Chat)
	if err := c.BodyParser(chat_to_delete); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	cur_chat := new(persistence.Chat)
	if err := persistence.DB.First(&cur_chat, chat_to_delete.ID).Error; err != nil {
		return c.Status(404).SendString(err.Error())
	}
	if cur_chat.ID != c.Locals("uid").(uint) {
		return c.SendStatus(401)
	}

	if err := persistence.DB.Delete(&chat_to_delete).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
}
