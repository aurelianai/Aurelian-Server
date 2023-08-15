package ahttp

import (
	"encoding/json"
	"net/http"
)

/*
Encodes data with json.Marshal and to writes to w with Status 200. This will escape html chars!

This can be returned in a ahttp.Handler
*/
func JSON(w http.ResponseWriter, data interface{}) (int, error) {
	encodedData, err := json.Marshal(data)

	if err != nil {
		return 500, err
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(encodedData)

	return 0, nil
}
