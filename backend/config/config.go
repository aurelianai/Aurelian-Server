package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Struct That Will be Used in Various Parts of the App to See what Model you have
type AelsConfig struct {
	Model     ModelConfig     `yaml:"model"`
	Inference InferenceConfig `yaml:"inference"`
}

type ModelConfig struct {
	System          string `yaml:"system"`
	UserPrefix      string `yaml:"userPrefix"`
	UserPostfix     string `yaml:"userPostfix"`
	ModelPrefix     string `yaml:"modelPrefix"`
	ModelPostfix    string `yaml:"modelPostfix"`
	ContextSize     int    `yaml:"contextSize"`
	ChatContextSize int    `yaml:"chatContextSize"`
}

type InferenceConfig struct {
	Endpoint     string `yaml:"endpoint"`
	Backend      string `yaml:"backend"`
	MaxNewTokens int    `yaml:"maxNewTokens"`
}

var Config = &AelsConfig{}

// Checks
func (c *AelsConfig) InitAndValidate(configFilePath string) error {
	fileContent, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	yaml.Unmarshal(fileContent, c)

	if c.Inference.Endpoint == "" {
		return errors.New("please pass an endpoint in your .yaml config file under inference:\n\ninference:\tendpoint: 'http://inf.com/generate'")
	}

	supportedBackends := `	1. text-generation-inference`

	if !strings.Contains(supportedBackends, c.Inference.Backend) {

		return fmt.Errorf("the inference backend '%s' is not supported\n\ntry one of these\n%s", c.Inference.Backend, supportedBackends)
	}

	if c.Model.ContextSize == 0 {
		return fmt.Errorf("context size must be greater than 0")
	}

	if c.Model.ChatContextSize == 0 {
		fmt.Printf("ChatContextSize defaulting to context size(%d tokens)\n", c.Model.ContextSize)
		c.Model.ChatContextSize = c.Model.ContextSize
	}

	if c.Model.ChatContextSize > c.Model.ContextSize {
		return fmt.Errorf("chat context (%d tokens) can't be larger than total context (%d tokens)", c.Model.ChatContextSize, c.Model.ContextSize)
	}

	if c.Inference.MaxNewTokens == 0 {
		c.Inference.MaxNewTokens = c.Model.ChatContextSize
	}

	return nil
}
