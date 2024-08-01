package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/omega-energia/code-review-copilot/ollama/constants"
)

type Loader interface {
	Load() (args, error)
}

type tinyLlamaLoader struct{}

func NewTinyLlamaLoader() *tinyLlamaLoader {
	return &tinyLlamaLoader{}
}

func (t *tinyLlamaLoader) Load() (args, error) {
	err := godotenv.Load()
	if err != nil {
		return args{}, err
	}

	return args{
		name:   os.Getenv(constants.ModelName),
		port:   os.Getenv(constants.ModelPort),
		prompt: os.Getenv(constants.ModelPrompt),
		temp:   os.Getenv(constants.ModelTemp),
	}, nil
}
