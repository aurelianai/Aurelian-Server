package logout

import (
	"AELS/ahttp"
	"net/http"
)

/*
Sets 'Auth' Cookie to ""
*/
func LogoutUser() ahttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {

		authCookie := http.Cookie{
			Name:  "Auth",
			Value: "",
		}
		http.SetCookie(w, &authCookie)

		return 0, nil
	}
}
