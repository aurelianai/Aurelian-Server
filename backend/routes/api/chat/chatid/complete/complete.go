package complete

import (
	a "AELS/ahttp"
	"AELS/lib/promptgen"
	m "AELS/middleware"
	"AELS/persistence"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"
)

type StreamResponse struct {
	Token Token `json:"token"`
}
type Token struct {
	Id      int     `json:"id"`
	Text    string  `json:"text"`
	LogProb float64 `json:"logprob"`
	Special bool    `json:"special"`
}

var STREAM_ENDPOINT = fmt.Sprintf("%s%s", os.Getenv("INFERENCE_BACKEND"), "/generate_stream")

/*
Complete through text/event-stream.
*/
func CompleteChat() a.Handler {

	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		chatid := r.Context().Value(m.ChatID{}).(uint64)

		var messages []persistence.Message
		err := persistence.DB.Where("chat_id = ?", chatid).Order("id ASC").Find(&messages).Error

		if err != nil {
			return 500, err
		}

		prompt, err := promptgen.GeneratePrompt(messages)

		if err != nil {
			return 500, err
		}

		reqBody, err := a.Marshal(a.Map{"inputs": prompt, "max_new_tokens": 50})

		if err != nil {
			return 500, err
		}

		req, err := http.NewRequest(http.MethodPost, STREAM_ENDPOINT, bytes.NewReader(reqBody))

		if err != nil {
			return 500, err
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			return 500, err
		}

		w.Header().Add("Content-Type", "text/event-stream")

		flusher := w.(http.Flusher)
		errors := make(chan error)
		completed := make(chan struct{})      // Signals Generation successfully completed
		canceled := make(chan struct{})       // Signals that Client Hung up and generation should stop
		gracefullyShutdown := make(chan bool) // Signals that generation has been shutdown gracefully

		go func() {

			scanner := bufio.NewReader(resp.Body)

			modelResponse := persistence.Message{Role: "MODEL", Content: "", ChatID: chatid}
			persistence.DB.Create(&modelResponse)

			for {
				select {

				case <-canceled:
					gracefullyShutdown <- true // Cannot close from within routine!
					return

				default:
					bin, err := scanner.ReadBytes('\n')

					if err != nil && err != io.EOF {
						errors <- err
						break
					}

					if len(bin) == 1 {
						continue
					}

					var streamResponse StreamResponse
					msgPrefix := []byte("data:")
					if bytes.HasPrefix(bin, msgPrefix) {
						jsonBin := bytes.TrimPrefix(bin, msgPrefix)
						if err := json.Unmarshal(jsonBin, &streamResponse); err != nil {
							errors <- err
							break
						}
					} else {
						continue
					}

					// TODO Batch This
					persistence.DB.Model(&persistence.Message{}).
						Where("id = ?", modelResponse.ID).
						Update("content", gorm.Expr("CONCAT(content, ?)", streamResponse.Token.Text))

					fmt.Printf("Recieved token: %s\n", streamResponse.Token.Text)

					if err == io.EOF {
						break
					}

					_, err = fmt.Fprintf(w, "%s\n", bin)
					if err != nil {
						errors <- err
						break
					}
					flusher.Flush()

					if streamResponse.Token.Special {
						close(completed)
						return
					}
				}
			}
		}()

		for {
			select {
			case <-r.Context().Done():
				close(canceled)
				fmt.Println("Client Hung Up Connection, awaiting generation shutdown")
				select {
				case <-gracefullyShutdown:
					fmt.Println("Generation gracefully stopped")
				case <-time.After(30 * time.Second):
					fmt.Println("Generation stop timed out after 5 seconds forcing panic")
				}
				return 0, nil

			case err := <-errors:
				return 500, fmt.Errorf("error occurred during generation: %s", err.Error())

			case <-completed:
				fmt.Println("Finished Generation Successfully")
				return 0, nil
			}
		}
	}
}
