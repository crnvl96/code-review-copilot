package spec

import (
	"log"
	"os"

	"github.com/omega-energia/code-review-copilot/internal/env"
	"github.com/omega-energia/code-review-copilot/pkg/validation"
)

type SpecInterface interface {
	FromEnv() SpecConstants
}

type Spec struct {
	envLoader env.EnvLoaderInterface
	validator validation.ValidatorInterface
}

func NewSpec(validator validation.ValidatorInterface, envLoader env.EnvLoaderInterface) *Spec {
	return &Spec{validator: validator, envLoader: envLoader}
}

type SpecConstants struct {
	AiModelName   string
	AiPort        string
	AiPrompt      string
	AiTemperature string
	AiBaseUrl     string
}

func (e *Spec) FromEnv() SpecConstants {
	e.envLoader.Load()

	const (
		aiModelName    = "AI_MODEL_NAME"
		aiPort         = "AI_PORT"
		aiPrompt       = "AI_PROMPT"
		aiTemperature  = "AI_TEMPERATURE"
		aiPrintByChunk = "AI_PRINT_BY_CHUNK"
		aiBaseURL      = "AI_BASE_URL"
	)

	modelName, err := e.validator.ModelName(os.Getenv(aiModelName))
	if err != nil {
		log.Fatal(err)
	}

	port, err := e.validator.ModelPort(os.Getenv(aiPort))
	if err != nil {
		log.Fatal(err)
	}

	prompt, err := e.validator.ModelPrompt(os.Getenv(aiPrompt))
	if err != nil {
		log.Fatal(err)
	}

	temp, err := e.validator.ModelTemperature(os.Getenv(aiTemperature))
	if err != nil {
		log.Fatal(err)
	}

	url, err := e.validator.ModelBaseURL(os.Getenv(aiBaseURL))
	if err != nil {
		log.Fatal(err)
	}

	return SpecConstants{
		AiModelName:   modelName,
		AiPort:        port,
		AiPrompt:      prompt,
		AiTemperature: temp,
		AiBaseUrl:     url,
	}
}
