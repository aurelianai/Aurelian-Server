package login

import (
	"AELS/ahttp"
	"AELS/middleware"
	"AELS/persistence"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type LoginPayload struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

/*
Email and Password no match --> 404

Email and Password match --> Assigns Auth cookie --> 200

other err --> 500
*/
func LoginUser() ahttp.Handler {

	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		loginPayload := new(LoginPayload)
		if err := ahttp.ParseBody(r, loginPayload); err != nil {
			return 500, err
		}

		user := new(persistence.User)
		err := persistence.DB.Where("email = ? AND password = ?",
			&loginPayload.Email,
			&loginPayload.Pass).
			First(user).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 401, errors.New("user not found")
		} else if err != nil {
			return 500, err
		}

		if err := middleware.AssignNewCookie(user.ID, w); err != nil {
			return 500, err
		}

		w.WriteHeader(200)
		return 0, nil
	}
}
