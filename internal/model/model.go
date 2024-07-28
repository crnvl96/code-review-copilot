package model

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/omega-energia/code-review-copilot/pkg/env"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func Spec() *ollama.LLM {
	e := env.Retrieve()
	serverURL := e.AiBaseUrl + e.AiPort

	llm, err := ollama.New(
		ollama.WithModel(e.AiModelName),
		ollama.WithServerURL(serverURL),
	)
	if err != nil {
		log.Fatal(err)
	}

	return llm
}

func Temperature() float64 {
	bits := 64
	e := env.Retrieve()

	temp, err := strconv.ParseFloat(e.AiTemperature, bits)
	if err != nil {
		log.Fatal(err)
	}

	return temp
}

func StreamingFunc(ctx context.Context, chunk []byte) error {
	fmt.Print(string(chunk))

	return nil
}

func Generate(
	ctx context.Context,
	llm *ollama.LLM,
	prompt string,
	temp float64,
	streamingFunc func(ctx context.Context, chunk []byte) error,
) {
	res, err := llms.GenerateFromSinglePrompt(
		ctx, llm, prompt, llms.WithTemperature(temp), llms.WithStreamingFunc(streamingFunc),
	)
	if err != nil {
		log.Fatal(err)
	}

	_ = res
}
