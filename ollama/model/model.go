package model

import (
	"context"
	"fmt"

	"github.com/omega-energia/code-review-copilot/ollama/config"
	"github.com/omega-energia/code-review-copilot/ollama/constants"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type model interface {
	Run()
}

type tinyLlama struct {
	config config.Params
}

func NewTinyLlama(config config.Params) *tinyLlama {
	return &tinyLlama{config: config}
}

func (t *tinyLlama) Run() error {
	URL := constants.ModelBaseUrl + t.config.Port

	ctx := context.Background()

	llm, err := ollama.New(ollama.WithModel(t.config.Name), ollama.WithServerURL(URL))
	if err != nil {
		return (err)
	}

	res, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		t.config.Prompt,
		llms.WithTemperature(t.config.Temp),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	if err != nil {
		return err
	}

	_ = res

	return nil
}
