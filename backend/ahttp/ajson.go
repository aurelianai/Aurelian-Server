package ahttp

import (
	"bytes"
	"encoding/json"
	"net/http"
)

/*
Same behavior as json.Marshal, but doesn't escape HTML
*/
func Marshal(body interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(body)
	return buffer.Bytes(), err
}

/*
Reads request body and uses json.Unmarshal to bind it the passed struct pointer

returns any errors, including 500 if syntax is invalid
*/
func ParseBody(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil
}

type Map map[string]interface{}
