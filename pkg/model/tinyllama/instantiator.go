package tinyllama

import (
	"context"
	"fmt"

	"github.com/omega-energia/code-review-copilot/pkg/model"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type modelInstanciator struct {
	parser model.Parser
}

func newModelInstanciator(parser model.Parser) *modelInstanciator {
	return &modelInstanciator{parser: parser}
}

func (t *modelInstanciator) Run() error {
	config, err := t.parser.Parse()
	if err != nil {
		return err
	}

	url := model.ModelBaseUrl + config.Port

	ctx := context.Background()

	llm, err := ollama.New(ollama.WithModel(config.Name), ollama.WithServerURL(url))
	if err != nil {
		return (err)
	}

	res, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		config.Prompt,
		llms.WithTemperature(config.Temp),
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
