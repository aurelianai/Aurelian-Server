package promptgen

import (
	"AELS/persistence"
	"fmt"
	"strings"

	"github.com/sugarme/tokenizer"
	"github.com/sugarme/tokenizer/pretrained"
)

/*
Generates prompt based on OpenAssistant prompt template

<prompter>{PROMPT}<endoftext><assistant>{RESPONSE}

Tokenization is done on the fly to be flexible to changes in the future
*/
func GeneratePrompt(messages []persistence.Message) (string, error) {
	configFile, err := tokenizer.CachedPath("OpenAssistant/falcon-7b-sft-mix-2000", "tokenizer.json")
	if err != nil {
		fmt.Printf("FATAL - Loading tokenizer 'tiiuae/falcon-7b' '%s'", err.Error())
		return "", err
	}

	tk, err := pretrained.FromFile(configFile)
	if err != nil {
		fmt.Printf("FATAL - Loading tokenizer 'tiiuae/falcon-7b' '%s'", err.Error())
		return "", err
	}

	promptBuilder := strings.Builder{}
	tokenCount := 0

	for _, message := range messages {
		en, err := tk.EncodeSingle(message.Content)
		if err != nil {
			return "", err
		}
		messageTokenSize := len(en.Tokens)
		if messageTokenSize+tokenCount <= 2047 {
			if message.Role == "USER" {
				_, err := promptBuilder.WriteString("<|prompter|>")
				if err != nil {
					return "", err
				}
			} else {
				_, err := promptBuilder.WriteString("<|assistant|>")
				if err != nil {
					return "", err
				}
			}
			_, err = promptBuilder.WriteString(message.Content)
			if err != nil {
				return "", err
			}
			_, err = promptBuilder.WriteString("<|endoftext|>")
			if err != nil {
				return "", err
			}
			tokenCount += messageTokenSize + 2
		} else {
			break
		}
	}
	promptBuilder.WriteString("<|assistant|>")

	return promptBuilder.String(), nil
}
