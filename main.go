package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/omega-energia/code-review-copilot/pkg/env"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	e := env.Retrieve()

	llm, err := ollama.New(
		ollama.WithModel(e.AiModelName),
		ollama.WithServerURL(e.AiBaseUrl+e.AiPort),
	)
	if err != nil {
		log.Fatal(err)
	}

	temp, err := strconv.ParseFloat(e.AiTemperature, 64)
	if err != nil {
		log.Fatal(err)
	}

	shouldPrintByChunk, err := strconv.ParseBool(e.AiPrintByChunk)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	_, err = llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		e.AiPrompt+"\n",
		llms.WithTemperature(temp),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))

			if shouldPrintByChunk {
				fmt.Print("\n")
			}

			return nil
		}),
	)

	if !shouldPrintByChunk {
		fmt.Print("\n")
	}

	if err != nil {
		log.Fatal(err)
	}
}
