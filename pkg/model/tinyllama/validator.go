package tinyllama

import (
	"errors"
	"log"
	"strconv"

	"github.com/omega-energia/code-review-copilot/pkg/model"
)

type tinyLlamaArgsValidator struct {
	loader model.Loader
}

func newTinyLlamaArgsValidator(loader model.Loader) *tinyLlamaArgsValidator {
	return &tinyLlamaArgsValidator{
		loader: loader,
	}
}

func (t *tinyLlamaArgsValidator) Validate() (model.ModelArgs, error) {
	args, err := t.loader.Load()
	if err != nil {
		return model.ModelArgs{}, err
	}

	name, err := validateName(args.Name)
	if err != nil {
		return model.ModelArgs{}, err
	}

	port, err := validatePort(args.Port)
	if err != nil {
		return model.ModelArgs{}, err
	}

	prompt, err := validatePrompt(args.Prompt)
	if err != nil {
		return model.ModelArgs{}, err
	}

	temp, err := validateTemp(args.Temp)
	if err != nil {
		return model.ModelArgs{}, err
	}

	return model.ModelArgs{
		Name:   name,
		Port:   port,
		Prompt: prompt,
		Temp:   temp,
	}, nil
}

func validateName(i string) (string, error) {
	if i == "" {
		return "", errors.New(model.ErrEmptyModelName)
	}

	if i != model.ModelDefaultName {
		return "", errors.New(model.ErrInvalidModelName)
	}

	return i, nil
}

func validatePort(i string) (string, error) {
	if i == "" {
		return "", errors.New(model.ErrEmptyModelPort)
	}

	if i == model.ModelDefaultPort {
		return "", errors.New(model.ErrInvalidModelPort)
	}

	return i, nil
}

func validatePrompt(i string) (string, error) {
	if i == "" {
		return "", errors.New(model.ErrEmptyModelPrompt)
	}

	if length, _ := strconv.Atoi(model.ModelMaxPromptLength); len(i) > length {
		return "", errors.New(model.ErrInvalidModelPrompt)
	}

	return i, nil
}

func validateTemp(i string) (string, error) {
	if i == "" {
		return "", errors.New(model.ErrInvalidModelPrompt)
	}

	temp, err := strconv.ParseFloat(i, 64)
	if err != nil {
		log.Fatal(model.ErrInvalidModelTemp)
	}

	min, _ := strconv.ParseFloat(model.ModelMinimumTempValue, 64)
	max, _ := strconv.ParseFloat(model.ModelMaximunTempValue, 64)

	if temp < min || temp > max {
		return "", errors.New(model.ErrInvalidModelTemp)
	}

	return i, nil
}
