package user

import (
	"AELS/ahttp"
	"AELS/middleware"
	"AELS/persistence"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/sethvargo/go-password/password"
	"gorm.io/gorm"
)

/*
Retrieves User Information using ID from r.Context()
*/
func GetUser() ahttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		userid := r.Context().Value(middleware.UserID{}).(uint64)

		var u persistence.User
		err := persistence.DB.Where("id = ?", userid).First(&u).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, errors.New("user does not exist")
		} else if err != nil {
			return 500, err
		}

		return ahttp.JSON(w, u)
	}

}

/*
# Schema

{ email: string } --> { password: string } | error.

# Notes

Automatically generates a 16 character secure password, adds user into the database
*/
func CreateUser() ahttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		authorization := r.Header.Get("Authorization")

		if authorization != os.Getenv("USER_CREATE_ACCESS_KEY") {
			fmt.Printf("Got %s, but needed %s\n",
				authorization,
				os.Getenv("USER_CREATE_ACCESS_KEY"),
			)
			return 403, errors.New("unauthorized")
		}

		type UserSignUpPayload struct {
			Email string `json:"email"`
		}

		userSignUpPayload := new(UserSignUpPayload)
		if err := ahttp.ParseBody(r, &userSignUpPayload); err != nil {
			return 500, err
		}

		// TODO Change this to randomly pick how many symbols or numbers occur
		res, err := password.Generate(16, 4, 3, false, false)
		if err != nil {
			return 500, err
		}

		u := new(persistence.User)
		u.Email = userSignUpPayload.Email
		u.Password = res
		u.FirstName = strings.Split(u.Email, "@")[0]
		u.LastName = ""
		if err := persistence.DB.Create(u).Error; err != nil {
			return 500, err
		}

		return ahttp.JSON(w, ahttp.Map{"password": u.Password})
	}
}
