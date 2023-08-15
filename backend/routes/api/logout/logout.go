package logout

import (
	"AELS/ahttp"
	"net/http"
	"time"
)

/*
Sets 'Auth' Cookie to ""
*/
func LogoutUser() ahttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {

		authCookie := http.Cookie{
			Name:     "Auth",
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * -24),
			Value:    "",
			SameSite: http.SameSiteStrictMode,
			HttpOnly: true,
		}
		http.SetCookie(w, &authCookie)

		return 0, nil
	}
}
