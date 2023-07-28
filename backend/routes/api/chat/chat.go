package chat

import (
	"AELS/ahttp"
	"AELS/persistence"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ChatList() ahttp.Handler {

	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		userid := r.Context().Value("userid").(uint64)
		var chats []persistence.Chat
		if err := persistence.DB.Where("user_id = ?", userid).Order("id DESC").Find(&chats).Error; err != nil {
			return 500, fmt.Errorf("error retreiving chats for uid: %d, err: %s", userid, err.Error())
		}

		return ahttp.JSON(w, chats)
	}
}

func NewChat() ahttp.Handler {

	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		type NewChatPayload struct {
			Title string `json:"title"`
		}

		var newChatPayload NewChatPayload
		if err := ahttp.ParseBody(r, &newChatPayload); err != nil {
			return 500, err
		}

		new_chat := new(persistence.Chat)
		new_chat.UserID = r.Context().Value("userid").(uint64)
		new_chat.Title = newChatPayload.Title

		if err := persistence.DB.Create(&new_chat).Error; err != nil {
			return 500, err
		}

		return ahttp.JSON(w, new_chat)
	}
}

// Updates chat title, returns 200 on success
func UpdateChat() ahttp.Handler {

	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		type UpdateChatPayload struct {
			NewTitle string `json:"new_title" validate:"required,min=1,max=32"`
		}

		updateChat := new(UpdateChatPayload)
		if err := ahttp.ParseBody(r, updateChat); err != nil {
			return 500, err
		}

		validate := validator.New()
		if err := validate.Struct(updateChat); err != nil {
			return 400, err
		}

		chatid := r.Context().Value("chatid").(uint64)

		err := persistence.DB.Model(&persistence.Chat{}).
			Where("id = ?", chatid).
			Update("title", updateChat.NewTitle).Error

		if err != nil {
			return 500, err
		}

		w.WriteHeader(200)

		return 0, nil
	}
}

// Deletes by pkey, returns 200 on success
func DeleteChat() ahttp.Handler {

	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		chatid := r.Context().Value("chatid").(uint64)

		res := persistence.DB.Delete(&persistence.Chat{}, chatid)

		if res.Error != nil {
			return 500, fmt.Errorf("error deleting chat(%d) err:%s", chatid, res.Error.Error())
		}

		if res.RowsAffected == 0 {
			return 500, fmt.Errorf("tried to delete chat(%d) that does not exist", chatid)
		}

		w.WriteHeader(200)

		return 0, nil
	}
}
