package complete

import (
	a "AELS/ahttp"
	"AELS/lib/promptgen"
	m "AELS/middleware"
	"AELS/persistence"
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
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

		errs := make(chan error)
		completed := make(chan struct{})      // Signals Generation successfully completed
		canceled := make(chan struct{})       // Signals that Client Hung up and generation should stop
		gracefullyShutdown := make(chan bool) // Signals that generation has been shutdown gracefully

		go consumeAndWrite(resp, w, chatid, errs, gracefullyShutdown, canceled, completed)

		for {
			select {
			case <-r.Context().Done():
				close(canceled)
				fmt.Println("Client Hung Up Connection, awaiting generation shutdown")
				select {
				case <-gracefullyShutdown:
					fmt.Println("Generation gracefully stopped")
				case <-time.After(30 * time.Second):
					fmt.Println("Generation stop timed out after 30 seconds forcing panic")
				}
				return 0, nil

			case err := <-errs:
				return 500, fmt.Errorf("error occurred during generation: %s", err.Error())

			case <-completed:
				fmt.Println("Finished Generation Successfully")
				return 0, nil
			}
		}
	}
}

/*
Consumes events from resp and rewrites them to w after saving to db
*/
func consumeAndWrite(
	resp *http.Response,
	w http.ResponseWriter,
	chatid uint64,
	errs chan error,
	gracefullyShutdown chan bool,
	canceled chan struct{},
	completed chan struct{}) {

	flusher := w.(http.Flusher)

	scanner := bufio.NewReader(resp.Body)

	modelResponse := persistence.Message{Role: "MODEL", Content: "", ChatID: chatid}
	persistence.DB.Create(&modelResponse)

	for {
		select {

		case <-resp.Request.Context().Done():
			errs <- errors.New("inference backend hung up")

		case <-canceled:
			gracefullyShutdown <- true // Cannot close from within routine!
			return

		default:
			bin, err := scanner.ReadBytes('\n')

			if err != nil && err != io.EOF {
				errs <- err
				return
			}

			if len(bin) == 1 {
				continue
			}

			var streamResponse StreamResponse
			if err := json.Unmarshal(bytes.TrimPrefix(bin, []byte("data:")), &streamResponse); err != nil {
				errs <- err
				return
			}

			if streamResponse.Token.Special {
				close(completed)
				return
			}

			// TODO Batch This
			if err := persistence.DB.Model(&persistence.Message{}).Where("id = ?", modelResponse.ID).Update("content", gorm.Expr("content || ?", streamResponse.Token.Text)).Error; err != nil {
				errs <- err
				return
			}

			_, err = fmt.Fprintf(w, "%s\n", bin)

			if err != nil {
				errs <- err
				return
			}
			flusher.Flush()

		}
	}
}
