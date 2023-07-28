package middleware

import (
	"AELS/ahttp"
	"AELS/persistence"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

/*
chatid int is not given as a query or url parameter --> 400

chat does not exist --> 404

error retrieving chat --> 500

user does not own chat --> 401

user owns chat, chatid is set in context
*/
func ChatOwnership(next ahttp.Handler) ahttp.Handler {

	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		userid := r.Context().Value("userid").(uint64)
		chatid, err := parseChatid(r)
		if err != nil {
			return 400, err
		}

		var chat persistence.Chat
		err = persistence.DB.First(&chat, chatid).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, fmt.Errorf("chat with id(%d) does not exist", chatid)
		} else if err != nil {
			return 500, fmt.Errorf("error retrieving chat id(%d) '%s'", chatid, err.Error())
		}

		if chat.UserID != userid {
			return 401, fmt.Errorf("user id(%d) attempted to access chat id(%d) which they do not own", userid, chatid)
		}

		ctx := context.WithValue(r.Context(), ChatID{}, chatid)

		next.ServeHTTP(w, r.WithContext(ctx))

		return 0, nil
	}
}

type ChatID struct{}

/*
Checks for chatId that parses to uint64 in Query Params or URL Params

Checks for URL first
*/
func parseChatid(r *http.Request) (uint64, error) {
	vars := mux.Vars(r)
	query := r.URL.Query()

	if vars["chatid"] != "" {
		chatid, err := strconv.ParseUint(vars["chatid"], 10, 64)
		if err != nil {
			return 0, err
		}
		return chatid, nil
	} else if query.Get("chatid") != "" {
		chatid, err := strconv.ParseUint(query.Get("chatid"), 10, 64)
		if err != nil {
			return 0, err
		}
		return chatid, nil
	} else {
		return 0, errors.New("no valid params")
	}

}
