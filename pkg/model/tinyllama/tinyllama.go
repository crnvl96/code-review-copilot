package tinyllama

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/omega-energia/code-review-copilot/pkg/model"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type tinyLlama struct{}

func NewTinyLlama() *tinyLlama {
	return &tinyLlama{}
}

func (t *tinyLlama) Run() error {
	params, err := generate()
	url := model.BaseUrl + params.Port

	ctx := context.Background()

	llm, err := ollama.New(ollama.WithModel(params.Name), ollama.WithServerURL(url))
	if err != nil {
		return (err)
	}

	res, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		params.Prompt,
		llms.WithTemperature(params.Temp),
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

func generate() (model.Params, error) {
	loadedArgs, err := load()
	if err != nil {
		return model.Params{}, err
	}

	validatedArgs, err := validate(loadedArgs)
	if err != nil {
		return model.Params{}, err
	}

	params, err := parse(validatedArgs)
	if err != nil {
		return model.Params{}, err
	}

	return params, nil
}

func load() (model.Args, error) {
	err := godotenv.Load()
	if err != nil {
		return model.Args{}, err
	}

	return model.Args{
		Name:   os.Getenv(model.Name),
		Port:   os.Getenv(model.Port),
		Prompt: os.Getenv(model.Prompt),
		Temp:   os.Getenv(model.Temp),
	}, nil
}

func validate(args model.Args) (model.Args, error) {
	name, err := validateName(args.Name)
	if err != nil {
		return model.Args{}, err
	}

	port, err := validatePort(args.Port)
	if err != nil {
		return model.Args{}, err
	}

	prompt, err := validatePrompt(args.Prompt)
	if err != nil {
		return model.Args{}, err
	}

	temp, err := validateTemp(args.Temp)
	if err != nil {
		return model.Args{}, err
	}

	return model.Args{
		Name:   name,
		Port:   port,
		Prompt: prompt,
		Temp:   temp,
	}, nil
}

func parse(args model.Args) (model.Params, error) {
	temp, err := strconv.ParseFloat(args.Temp, 64)
	if err != nil {
		return model.Params{}, err
	}

	return model.Params{
		Name:   args.Name,
		Port:   args.Port,
		Prompt: args.Prompt,
		Temp:   temp,
	}, nil
}

func validateName(i string) (string, error) {
	if i == "" {
		return "", errors.New(model.ErrEmptyName)
	}

	if i != model.DefaultName {
		return "", errors.New(model.ErrInvalidName)
	}

	return i, nil
}

func validatePort(i string) (string, error) {
	if i == "" {
		return "", errors.New(model.ErrEmptyPort)
	}

	if i == model.DefaultPort {
		return "", errors.New(model.ErrInvalidPort)
	}

	return i, nil
}

func validatePrompt(i string) (string, error) {
	if i == "" {
		return "", errors.New(model.ErrEmptyPrompt)
	}

	if length, _ := strconv.Atoi(model.MaxPromptLen); len(i) > length {
		return "", errors.New(model.ErrInvalidPrompt)
	}

	return i, nil
}

func validateTemp(i string) (string, error) {
	if i == "" {
		return "", errors.New(model.ErrInvalidTemp)
	}

	temp, err := strconv.ParseFloat(i, 64)
	if err != nil {
		log.Fatal(model.ErrInvalidTemp)
	}

	min, _ := strconv.ParseFloat(model.MinTempVal, 64)
	max, _ := strconv.ParseFloat(model.MaxTempVal, 64)

	if temp < min || temp > max {
		return "", errors.New(model.ErrInvalidTemp)
	}

	return i, nil
}
