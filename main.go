package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	llm, err := ollama.New(
		ollama.WithModel("tinyllama"),
		// ollama.WithServerURL("http://localhost:1111"), // 100% CPU based
		ollama.WithServerURL("http://localhost:2222"), // 34%/66% CPU/GPU based
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		"Give a brief description about the world cup 2022 finals\n",
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)

	fmt.Print("\n")

	if err != nil {
		log.Fatal(err)
	}

	_ = completion
}
