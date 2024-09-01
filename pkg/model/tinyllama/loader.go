package tinyllama

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/omega-energia/code-review-copilot/pkg/model"
)

type tinyLlamaArgsLoader struct{}

func newTinyLlamaArgsLoader() *tinyLlamaArgsLoader {
	return &tinyLlamaArgsLoader{}
}

func (a *tinyLlamaArgsLoader) Load() (model.ModelArgs, error) {
	err := godotenv.Load()
	if err != nil {
		return model.ModelArgs{}, err
	}

	return model.ModelArgs{
		Name:   os.Getenv(model.ModelName),
		Port:   os.Getenv(model.ModelPort),
		Prompt: os.Getenv(model.ModelPrompt),
		Temp:   os.Getenv(model.ModelTemp),
	}, nil
}
