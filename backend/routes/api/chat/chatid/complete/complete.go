package complete

import (
	a "AELS/ahttp"
	"AELS/config"
	"AELS/inference"
	m "AELS/middleware"
	"AELS/persistence"
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
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

		prompt, err := inference.GeneratePrompt(messages)
		if err != nil {
			return 500, err
		}

		backend := inference.SUPPORTED_BACKENDS[config.Config.Inference.Backend]
		resp, err := backend.StartStream(prompt)
		if err != nil {
			return 500, err
		}

		w.Header().Add("Content-Type", "text/event-stream")

		errs := make(chan error)              // Error Occurred during generation in the middle-end
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
	backend := inference.SUPPORTED_BACKENDS[config.Config.Inference.Backend]

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

			if string(bin) == "\n" {
				continue
			}

			event := backend.ParseEvent(bin)
			encodedEvent, err := a.Marshal(event)
			if err != nil {
				errs <- err
				return
			}

			_, err = fmt.Fprintf(w, "%s", encodedEvent)
			if err != nil {
				errs <- err
				return
			}
			flusher.Flush()

			if event.Err != nil || event.Last {
				close(completed)
				return
			}

			if err := persistence.DB.Model(&persistence.Message{}).Where("id = ?", modelResponse.ID).Update("content", gorm.Expr("content || ?", event.Delta)).Error; err != nil {
				errs <- err
				return
			}

		}
	}
}
