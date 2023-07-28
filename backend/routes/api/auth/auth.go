package auth

import (
	"AELS/ahttp"
	"net/http"
)

func CheckAuth() ahttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		w.WriteHeader(200)
		return 0, nil
	}
}
