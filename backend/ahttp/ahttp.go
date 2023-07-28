package ahttp

import "net/http"

type Handler func(http.ResponseWriter, *http.Request) (int, error)

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if code, err := fn(w, r); err != nil {
		http.Error(w, err.Error(), code)
	}
}
