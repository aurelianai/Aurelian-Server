package inference

import (
	"AELS/config"
	"AELS/persistence"
	"fmt"
	"strings"

	"github.com/sugarme/tokenizer"
	"github.com/sugarme/tokenizer/pretrained"
)

/*
Generates prompt string based on Given prompt template variables

{SYSTEM_PROMPT}{USER_PREFIX}{PROMPT}{USER_POSTFIX}{MODEL_PREFIX}{RESPONSE}{MODEL_POSTFIX}

Tokenization is done on the fly to be flexible to changes in the future
*/
func GeneratePrompt(messages []persistence.Message) (string, error) {
	// TODO Better Counting of Tokens, should be roughly the same for all models
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

	promptBuilder.WriteString(config.Config.Model.System)
	en, err := tk.EncodeSingle(config.Config.Model.System)
	if err != nil {
		return "", err
	}
	tokenCount += len(en.Tokens)

	for _, message := range messages {
		en, err := tk.EncodeSingle(message.Content)
		if err != nil {
			return "", err
		}
		messageTokenSize := len(en.Tokens)
		if messageTokenSize+tokenCount <= config.Config.Model.ChatContextSize {
			if message.Role == "USER" {
				_, err := promptBuilder.WriteString(config.Config.Model.UserPrefix)
				if err != nil {
					return "", err
				}
			} else {
				_, err := promptBuilder.WriteString(config.Config.Model.ModelPrefix)
				if err != nil {
					return "", err
				}
			}

			_, err = promptBuilder.WriteString(message.Content)
			if err != nil {
				return "", err
			}

			if message.Role == "USER" {
				_, err = promptBuilder.WriteString(config.Config.Model.UserPostfix)
				if err != nil {
					return "", err
				}
			} else {
				_, err = promptBuilder.WriteString(config.Config.Model.ModelPostfix)
				if err != nil {
					return "", err
				}
			}
			tokenCount += messageTokenSize + 2
		} else {
			break
		}
	}
	promptBuilder.WriteString(config.Config.Model.ModelPrefix)

	return promptBuilder.String(), nil
}
