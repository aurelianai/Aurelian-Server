package inference

import (
	a "AELS/ahttp"
	"AELS/config"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type TogetherAPI struct{}

func (*TogetherAPI) StartStream(prompt string) (*http.Response, error) {
	reqBody, err := a.Marshal(a.Map{
		"model":         "togethercomputer/llama-2-70b-chat",
		"prompt":        prompt,
		"max_tokens":    config.Config.Inference.MaxNewTokens,
		"stream_tokens": true,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, config.Config.Inference.Endpoint, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	if os.Getenv("TOGETHER_API_KEY") == "" {
		return nil, errors.New("set TOGETHER_API_KEY environment variable")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TOGETHER_API_KEY")))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (*TogetherAPI) ParseEvent(newLine []byte) InferenceUpdate {
	data := bytes.Trim(
		bytes.TrimPrefix(newLine, []byte("data:")),
		"\n ",
	)

	if string(data) == "[DONE]" {
		return InferenceUpdate{Delta: "", Err: nil, Last: true}
	}

	type Choice struct {
		Text string
	}
	type StreamResponse struct {
		Choices []Choice `json:"choices"`
	}

	var streamResponse StreamResponse
	if err := json.Unmarshal(data, &streamResponse); err != nil {
		fmt.Printf("Error occured deserializing streamResponse: %s", err.Error())
		return InferenceUpdate{Delta: "", Err: err, Last: false}
	}
	if len(streamResponse.Choices) == 0 {
		return InferenceUpdate{Delta: "", Err: errors.New("nothing returned"), Last: false}
	}

	return InferenceUpdate{Delta: streamResponse.Choices[0].Text, Err: nil, Last: false}
}
