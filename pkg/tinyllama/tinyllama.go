package tinyllama

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

const (
	baseUrl          = "http://localhost:"
	name             = "AI_MODEL_NAME"
	port             = "AI_PORT"
	prompt           = "AI_PROMPT"
	temp             = "AI_TEMPERATURE"
	defaultPort      = "11434"
	defaultName      = "tinyllama"
	maxPromptLen     = "120"
	minTempVal       = "0.1"
	maxTempVal       = "1.0"
	errEmptyName     = "value of \"AI_MODEL_NAME\" must not be empty"
	errInvalidName   = "value of \"AI_MODEL_NAME\" must be \"tinyllama\""
	errEmptyPort     = "value of \"AI_PORT\" must not be empty"
	errInvalidPort   = "value of \"AI_PORT\" must not be " + defaultPort
	errEmptyPrompt   = "value of \"AI_PROMPT\" must not be empty"
	errInvalidPrompt = "length of \"AI_PROMPT\" be less than " + maxPromptLen
	errEmptyTemp     = "value of \"AI_TEMPERATURE\" must not be empty"
	errInvalidTemp   = "value of \"AI_TEMPERATURE\" must be a valid number between " + minTempVal + " and " + maxTempVal
)

type modelParams struct {
	name   string
	port   string
	prompt string
	temp   float64
}

type modelArgs struct {
	name   string
	port   string
	prompt string
	temp   string
}

func Run(data string) (string, error) {
	params, err := generate()
	url := baseUrl + port

	ctx := context.Background()

	llm, err := ollama.New(ollama.WithModel(params.name), ollama.WithServerURL(url))
	if err != nil {
		return "", err
	}

	prompt := params.prompt + data

	res, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		prompt,
		llms.WithTemperature(params.temp),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	if err != nil {
		return "", err
	}

	return res, nil
}

func generate() (modelParams, error) {
	loadedArgs, err := load()
	if err != nil {
		return modelParams{}, err
	}

	validatedArgs, err := validate(loadedArgs)
	if err != nil {
		return modelParams{}, err
	}

	params, err := parse(validatedArgs)
	if err != nil {
		return modelParams{}, err
	}

	return params, nil
}

func load() (modelArgs, error) {
	err := godotenv.Load()
	if err != nil {
		return modelArgs{}, err
	}

	return modelArgs{
		name:   os.Getenv(name),
		port:   os.Getenv(port),
		prompt: os.Getenv(prompt),
		temp:   os.Getenv(temp),
	}, nil
}

func validate(args modelArgs) (modelArgs, error) {
	name, err := validateName(args.name)
	if err != nil {
		return modelArgs{}, err
	}

	port, err := validatePort(args.port)
	if err != nil {
		return modelArgs{}, err
	}

	prompt, err := validatePrompt(args.prompt)
	if err != nil {
		return modelArgs{}, err
	}

	temp, err := validateTemp(args.temp)
	if err != nil {
		return modelArgs{}, err
	}

	return modelArgs{
		name:   name,
		port:   port,
		prompt: prompt,
		temp:   temp,
	}, nil
}

func parse(args modelArgs) (modelParams, error) {
	temp, err := strconv.ParseFloat(args.temp, 64)
	if err != nil {
		return modelParams{}, err
	}

	return modelParams{
		name:   args.name,
		port:   args.port,
		prompt: args.prompt,
		temp:   temp,
	}, nil
}

func validateName(i string) (string, error) {
	if i == "" {
		return "", errors.New(errEmptyName)
	}

	if i != defaultName {
		return "", errors.New(errInvalidName)
	}

	return i, nil
}

func validatePort(i string) (string, error) {
	if i == "" {
		return "", errors.New(errEmptyPort)
	}

	if i == defaultPort {
		return "", errors.New(errInvalidPort)
	}

	return i, nil
}

func validatePrompt(i string) (string, error) {
	if i == "" {
		return "", errors.New(errEmptyPrompt)
	}

	if length, _ := strconv.Atoi(maxPromptLen); len(i) > length {
		return "", errors.New(errInvalidPrompt)
	}

	return i, nil
}

func validateTemp(i string) (string, error) {
	if i == "" {
		return "", errors.New(errInvalidTemp)
	}

	temp, err := strconv.ParseFloat(i, 64)
	if err != nil {
		log.Fatal(errInvalidTemp)
	}

	min, _ := strconv.ParseFloat(minTempVal, 64)
	max, _ := strconv.ParseFloat(maxTempVal, 64)

	if temp < min || temp > max {
		return "", errors.New(errInvalidTemp)
	}

	return i, nil
}
