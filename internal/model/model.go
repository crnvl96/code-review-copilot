package model

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/omega-energia/code-review-copilot/pkg/spec"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type ModelInterface interface {
	GenerateSpec() *ollama.LLM
	GenerateTemperature() float64
	StreamingFunc(ctx context.Context, chuk []byte) error
	GetResponse(
		ctx context.Context,
		llm *ollama.LLM,
		prompt string,
		temp float64,
		streamingFunc func(ctx context.Context, chunk []byte) error)
}

type Model struct {
	spec spec.SpecInterface
}

func NewModel(spec spec.SpecInterface) *Model {
	return &Model{spec: spec}
}

func (m *Model) GenerateSpec() *ollama.LLM {
	e := m.spec.FromEnv()
	URL := e.AiBaseUrl + e.AiPort

	llm, err := ollama.New(
		ollama.WithModel(e.AiModelName),
		ollama.WithServerURL(URL),
	)
	if err != nil {
		log.Fatal(err)
	}

	return llm
}

func (m *Model) GenerateTemperature() float64 {
	bits := 64

	temp, err := strconv.ParseFloat(m.spec.FromEnv().AiTemperature, bits)
	if err != nil {
		log.Fatal(err)
	}

	return temp
}

func (m *Model) StreamingFunc(ctx context.Context, chunk []byte) error {
	fmt.Print(string(chunk))

	return nil
}

func (m *Model) GetResponse(
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
