package inference

import (
	"net/http"
)

// Any New inference backend must define methods to format a request and to parse SSE outputs
type InferenceBackend interface {
	StartStream(prompt string) (*http.Response, error)
	ParseEvent([]byte) InferenceUpdate
}

// Universal SSE that's sent to the frontend
type InferenceUpdate struct {
	Delta string `json:"delta"`
	Err   error  `json:"err"`
	Last  bool   `json:"last"`
}

var SUPPORTED_BACKENDS = map[string]InferenceBackend{
	"text-generation-inference": &TextGenerationInference{},
}
