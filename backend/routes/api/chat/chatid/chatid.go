package chatid

import (
	"AELS/ahttp"
	"AELS/persistence"
	"fmt"
	"net/http"
)

/*
Error getting messages --> 500

Otherwise --> JSON
*/
func ListMessages() ahttp.Handler {

	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		chatid := r.Context().Value("chatid").(uint64)

		var messages []persistence.Message
		err := persistence.DB.Where("chat_id = ?", chatid).Order("id ASC").Find(&messages).Error

		if err != nil {
			return 500, fmt.Errorf("error Collecting Messages: %s", err.Error())
		}

		return ahttp.JSON(w, messages)
	}
}

/*
Error parsing body --> 400

Error with db interaction --> 500

Success --> 200 and JSON of Message
*/
func NewMessage() ahttp.Handler {

	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		chatid := r.Context().Value("chatid").(uint64)

		var new_message = new(persistence.Message)
		err := ahttp.ParseBody(r, &new_message)

		if err != nil {
			return 400, fmt.Errorf("error Parsing message body: %s", err.Error())
		}

		new_message.ChatID = chatid
		err = persistence.DB.Create(&new_message).Error

		if err != nil {
			return 500, fmt.Errorf("error creating new message: %s", err.Error())
		}

		return ahttp.JSON(w, new_message)
	}
}
