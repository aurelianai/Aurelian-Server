package complete

import (
	"AELS/ahttp"
	"net/http"
)

type TextGenPayload struct {
	Inputs     string            `json:"inputs"`
	Parameters TextGenParameters `json:"parameters"`
}

type TextGenParameters struct {
	MaxNewTokens uint `json:"max_new_tokens"`
}

type TextGenResponse struct {
	GeneratedText string `json:"generated_text"`
}

/*
Complete through text/event-stream.
*/
func CompleteChat() ahttp.Handler {

	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		return 0, nil
	}
}
