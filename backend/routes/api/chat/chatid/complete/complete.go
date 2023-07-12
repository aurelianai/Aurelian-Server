package complete

import (
	"AELS/lib/promptgen"
	"AELS/persistence"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
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
Most complicated request handler in the repo.
Returns a handler becuase it uses fasthttp to make request to tgi
Creates new chat, saves it, and returns json.
*/
func CompleteChat() fiber.Handler {
	client := fasthttp.Client{}

	return func(c *fiber.Ctx) error {
		chatId := c.Locals("chatid").(uint64)

		var messages []persistence.Message
		persistence.DB.Where("chat_id = ?", chatId).Find(&messages)

		p, err := promptgen.GeneratePrompt(messages)
		if err != nil {
			return err
		}

		body, err := encodeTextGenPayload(TextGenPayload{
			Inputs: p,
			Parameters: TextGenParameters{
				MaxNewTokens: 30,
			},
		})
		if err != nil {
			return err
		}

		req := fasthttp.AcquireRequest()
		res := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(res)

		req.SetRequestURI(os.Getenv("INFERENCE_BACKEND") + "/generate")
		req.Header.SetMethod("POST")
		req.Header.SetContentType("application/json")
		req.SetBody(body)

		if err := client.DoTimeout(req, res, 20*time.Second); err != nil {
			return err
		}

		if res.StatusCode() != 200 {
			fmt.Printf("Recieved code %d from inference backend\n", res.StatusCode())
			return errors.New("recieved status code from inference backend")
		}

		c.Set("Content-Type", "application/json")

		var textGenResponse TextGenResponse
		json.Unmarshal(res.Body(), &textGenResponse)

		responseMessage := new(persistence.Message)
		responseMessage.Role = "MODEL"
		responseMessage.Content = textGenResponse.GeneratedText
		responseMessage.ChatID = chatId
		if err := persistence.DB.Create(&responseMessage).Error; err != nil {
			err := fmt.Sprintf("error creating new message: '%s'", err.Error())
			fmt.Println(err)
			return c.Status(500).SendString(err)
		}

		return c.JSON(responseMessage)
	}
}

// Helper for Encoding
func encodeTextGenPayload(body interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(body)
	return buffer.Bytes(), err
}
