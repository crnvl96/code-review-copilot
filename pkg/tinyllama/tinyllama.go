package tinyllama

import (
	"context"
	"fmt"

	"github.com/crnvl96/code-review-copilot/internal/constants"
	"github.com/crnvl96/code-review-copilot/pkg/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GenerateModel(
	modelConfig config.ModelConfig,
) (func(modelPromptDetails string) (string, error), error) {
	config := modelConfig.Config

	url := constants.LlmContainerBaseUrl + config.Port
	baseModel := ollama.WithModel(config.Name)
	server := ollama.WithServerURL(url)

	llm, err := ollama.New(baseModel, server)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	return func(modelPromptDetails string) (string, error) {
		res, err := llms.GenerateFromSinglePrompt(
			ctx,
			llm,
			fmt.Sprintf("%s\n%s", config.Prompt, modelPromptDetails),
			llms.WithTemperature(config.Temp),
			llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				fmt.Print(string(chunk))
				return nil
			}),
		)
		if err != nil {
			return "", err
		}

		return res, nil
	}, nil
}
