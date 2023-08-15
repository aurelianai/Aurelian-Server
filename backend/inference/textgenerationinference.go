package inference

import (
	a "AELS/ahttp"
	"AELS/config"
	"bytes"
	"encoding/json"
	"net/http"
)

type TextGenerationInference struct{}

func (*TextGenerationInference) StartStream(prompt string) (*http.Response, error) {
	reqBody, err := a.Marshal(a.Map{"inputs": prompt})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, config.Config.Inference.Endpoint, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (*TextGenerationInference) ParseEvent(newLine []byte) InferenceUpdate {
	type Token struct {
		Id      int     `json:"id"`
		Text    string  `json:"text"`
		LogProb float64 `json:"logprob"`
		Special bool    `json:"special"`
	}

	type StreamResponse struct {
		Token Token `json:"token"`
	}

	var streamResponse StreamResponse
	if err := json.Unmarshal(bytes.TrimPrefix(newLine, []byte("data:")), &streamResponse); err != nil {
		return InferenceUpdate{Delta: "", Err: err, Last: false}
	}

	return InferenceUpdate{Delta: streamResponse.Token.Text, Err: nil, Last: streamResponse.Token.Special}
}
