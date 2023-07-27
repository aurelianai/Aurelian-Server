package promptgen

import (
	"AELS/persistence"
	"fmt"
	"strings"

	"github.com/sugarme/tokenizer"
	"github.com/sugarme/tokenizer/pretrained"
)

var (
	SYSTEM_PROMPT string = "### human: Interact in conversation to the best of your ability, please be concise, logical, intelligent and coherent.\n\n### response: Sure! sounds good.\n\n"
	USER_PREFIX   string = "### human: "
	USER_POSTFIX  string = "\n\n"
	MODEL_PREFIX  string = "### response: "
	MODEL_POSTFIX string = "\n\n"
)

/*
Generates prompt string based on Given prompt template variables

{SYSTEM_PROMPT}{USER_PREFIX}{PROMPT}{USER_POSTFIX}{MODEL_PREFIX}{RESPONSE}{MODEL_POSTFIX}

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

	en, err := tk.EncodeSingle(SYSTEM_PROMPT)
	if err != nil {
		return "", err
	}
	promptBuilder.WriteString(SYSTEM_PROMPT)
	tokenCount += len(en.Tokens)

	for _, message := range messages {
		en, err := tk.EncodeSingle(message.Content)
		if err != nil {
			return "", err
		}
		messageTokenSize := len(en.Tokens)
		if messageTokenSize+tokenCount <= 4096 {
			if message.Role == "USER" {
				_, err := promptBuilder.WriteString(USER_PREFIX)
				if err != nil {
					return "", err
				}
			} else {
				_, err := promptBuilder.WriteString(MODEL_PREFIX)
				if err != nil {
					return "", err
				}
			}

			_, err = promptBuilder.WriteString(message.Content)
			if err != nil {
				return "", err
			}

			if message.Role == "USER" {
				_, err = promptBuilder.WriteString(USER_POSTFIX)
				if err != nil {
					return "", err
				}
			} else {
				_, err = promptBuilder.WriteString(MODEL_POSTFIX)
				if err != nil {
					return "", err
				}
			}
			tokenCount += messageTokenSize + 2
		} else {
			break
		}
	}
	promptBuilder.WriteString(MODEL_PREFIX)

	return promptBuilder.String(), nil
}
