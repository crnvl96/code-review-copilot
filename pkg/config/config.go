package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/crnvl96/code-review-copilot/internal/constants"
)

type Params struct {
	Name   string
	Port   string
	Prompt string
	Temp   float64
}

type Args struct {
	Name   string
	Port   string
	Prompt string
	Temp   string
}

type ModelConfig struct {
	Config Params
}

func GenerateConfig() (ModelConfig, error) {
	rawArgs := load()

	validatedArgs, err := validate(rawArgs)
	if err != nil {
		return ModelConfig{}, err
	}

	parsedArgs, err := parse(validatedArgs)
	if err != nil {
		return ModelConfig{}, err
	}

	return ModelConfig{
		Config: Params{
			Name:   parsedArgs.Name,
			Port:   parsedArgs.Port,
			Prompt: parsedArgs.Prompt,
			Temp:   parsedArgs.Temp,
		},
	}, nil
}

func load() Args {
	return Args{
		Name:   os.Getenv(constants.LlmName),
		Port:   os.Getenv(constants.LlmPort),
		Prompt: os.Getenv(constants.LlmPrompt),
		Temp:   os.Getenv(constants.LlmTemp),
	}
}

func validate(args Args) (Args, error) {
	fmt.Println(args)

	name, err := validateName(args.Name)
	if err != nil {
		return Args{}, err
	}

	port, err := validatePort(args.Port)
	if err != nil {
		return Args{}, err
	}

	prompt, err := validatePrompt(args.Prompt)
	if err != nil {
		return Args{}, err
	}

	temp, err := validateTemp(args.Temp)
	if err != nil {
		return Args{}, err
	}

	return Args{
		Name:   name,
		Port:   port,
		Prompt: prompt,
		Temp:   temp,
	}, nil
}

func parse(args Args) (Params, error) {
	temp, err := strconv.ParseFloat(args.Temp, 64)
	if err != nil {
		return Params{}, err
	}

	return Params{
		Name:   args.Name,
		Port:   args.Port,
		Prompt: args.Prompt,
		Temp:   temp,
	}, nil
}

func validateName(i string) (string, error) {
	if i == "" {
		return "", errors.New(constants.ErrEmptyLlmName)
	}

	if i != constants.LlmName {
		return "", errors.New(constants.ErrInvalidLlmName)
	}

	return i, nil
}

func validatePort(i string) (string, error) {
	if i == "" {
		return "", errors.New(constants.ErrEmptyLlmPort)
	}

	if i == constants.LlmDefaultPort {
		return "", errors.New(constants.ErrInvalidLlmPort)
	}

	return i, nil
}

func validatePrompt(i string) (string, error) {
	if i == "" {
		return "", errors.New(constants.ErrEmptyLlmPrompt)
	}

	if length, _ := strconv.Atoi(constants.LlmMaxPromptLen); len(i) > length {
		return "", errors.New(constants.ErrInvalidLlmPrompt)
	}

	return i, nil
}

func validateTemp(i string) (string, error) {
	if i == "" {
		return "", errors.New(constants.ErrEmptyLlmTemp)
	}

	temp, err := strconv.ParseFloat(i, 64)
	if err != nil {
		log.Fatal(constants.ErrInvalidLlmTemp)
	}

	min, _ := strconv.ParseFloat(constants.LlmMinTempVal, 64)
	max, _ := strconv.ParseFloat(constants.LlmMaxTempVal, 64)

	if temp < min || temp > max {
		return "", errors.New(constants.ErrInvalidLlmTemp)
	}

	return i, nil
}
