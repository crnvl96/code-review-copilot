package config

import (
	"errors"
	"log"
	"strconv"

	"github.com/omega-energia/code-review-copilot/ollama/constants"
)

type Validator interface {
	Validate(args) (args, error)
}

type tinyLlamaValidator struct{}

func NewTinyLlamaValidator() *tinyLlamaValidator {
	return &tinyLlamaValidator{}
}

func (t *tinyLlamaValidator) Validate(a args) (args, error) {
	name, err := validateName(a.name)
	if err != nil {
		return args{}, err
	}

	port, err := validatePort(a.port)
	if err != nil {
		return args{}, err
	}

	prompt, err := validatePrompt(a.prompt)
	if err != nil {
		return args{}, err
	}

	temp, err := validateTemp(a.temp)
	if err != nil {
		return args{}, err
	}

	return args{
		name:   name,
		port:   port,
		prompt: prompt,
		temp:   temp,
	}, nil
}

func validateName(i string) (string, error) {
	if i == "" {
		return "", errors.New(constants.ErrEmptyModelName)
	}

	if i != constants.ModelDefaultName {
		return "", errors.New(constants.ErrInvalidModelName)
	}

	return i, nil
}

func validatePort(i string) (string, error) {
	if i == "" {
		return "", errors.New(constants.ErrEmptyModelPort)
	}

	if i == constants.ModelDefaultPort {
		return "", errors.New(constants.ErrInvalidModelPort)
	}

	return i, nil
}

func validatePrompt(i string) (string, error) {
	if i == "" {
		return "", errors.New(constants.ErrEmptyModelPrompt)
	}

	if length, _ := strconv.Atoi(constants.ModelMaxPromptLength); len(i) > length {
		return "", errors.New(constants.ErrInvalidModelPrompt)
	}

	return i, nil
}

func validateTemp(i string) (string, error) {
	if i == "" {
		return "", errors.New(constants.ErrInvalidModelPrompt)
	}

	temp, err := strconv.ParseFloat(i, 64)
	if err != nil {
		log.Fatal(constants.ErrInvalidModelTemp)
	}

	min, _ := strconv.ParseFloat(constants.ModelMinimumTempValue, 64)
	max, _ := strconv.ParseFloat(constants.ModelMaximunTempValue, 64)

	if temp < min || temp > max {
		return "", errors.New(constants.ErrInvalidModelTemp)
	}

	return i, nil
}
