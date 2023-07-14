package chat

import (
	"AELS/persistence"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ChatList(c *fiber.Ctx) error {
	uid := c.Locals("uid").(uint64)
	var chats []persistence.Chat
	if err := persistence.DB.Where("user_id = ?", uid).Order("id DESC").Find(&chats).Error; err != nil {
		err := fmt.Sprintf("Error retreiving chats for uid: %d, err: %s", uid, err.Error())
		fmt.Println(err)
		return c.Status(500).SendString(err)
	}

	return c.JSON(chats)
}

func NewChat(c *fiber.Ctx) error {
	type NewChatPayload struct {
		Title string `json:"title"`
	}
	var newChatPayload NewChatPayload
	if err := c.BodyParser(&newChatPayload); err != nil {
		return err
	}

	new_chat := new(persistence.Chat)
	new_chat.UserID = c.Locals("uid").(uint64)
	new_chat.Title = newChatPayload.Title
	if err := persistence.DB.Create(&new_chat).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(new_chat)
}

// Updates chat title, returns 200 on success
func UpdateChat(c *fiber.Ctx) error {
	type UpdateChatPayload struct {
		NewTitle string `json:"new_title" validate:"required,min=1,max=32"`
	}

	update_chat := new(UpdateChatPayload)
	if err := c.BodyParser(update_chat); err != nil {
		fmt.Printf("%s", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(update_chat); err != nil {
		fmt.Println(err.Error())
		return c.Status(400).JSON(err)
	}

	err := persistence.DB.Model(&persistence.Chat{}).
		Where("id = ?", c.Locals("chatid").(uint64)).
		Update("title", update_chat.NewTitle).Error
	if err != nil {
		log.Printf("%s", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
}

// Deletes by pkey, returns 200 on success
func DeleteChat(c *fiber.Ctx) error {
	chatid := c.Locals("chatid").(uint64)
	res := persistence.DB.Delete(&persistence.Chat{}, chatid)
	if res.Error != nil {
		err := fmt.Sprintf("Error deleting chat(%d) err:%s", chatid, res.Error.Error())
		fmt.Println(err)
		return c.Status(500).SendString(err)
	}
	if res.RowsAffected == 0 {
		err := fmt.Sprintf("Tried to delete chat(%d) that does not exist", chatid)
		fmt.Println(err)
		return c.Status(404).SendString(err)
	}

	return c.SendStatus(200)
}
