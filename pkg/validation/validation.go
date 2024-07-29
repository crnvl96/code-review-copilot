package validation

import (
	"errors"
	"log"
	"reflect"
	"strconv"
)

type ValidatorInterface interface {
	ModelName(input string) (string, error)
	ModelPort(input string) (string, error)
	ModelPrompt(input string) (string, error)
	ModelTemperature(input string) (string, error)
	ModelBaseURL(input string) (string, error)
}

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ModelName(input string) (string, error) {
	models := map[string]string{
		"tinyllama": "tinyllama",
	}

	if _, exists := models[input]; !exists {
		return "", errors.New("AI_MODEL_NAME must be \"tinyllama\"")
	}

	return input, nil
}

func (v *Validator) ModelPort(input string) (string, error) {
	if input == "" {
		return "", errors.New("AI_PORT cannot be empty")
	}

	if reflect.TypeOf(input).Kind() != reflect.String {
		return "", errors.New("AI_PORT must be a string")
	}

	modelContainerPort := "11434"

	if input == modelContainerPort {
		return "", errors.New(
			modelContainerPort + " cannot be used as AI_PORT",
		)
	}

	return input, nil
}

func (v *Validator) ModelPrompt(input string) (string, error) {
	if input == "" {
		return "", errors.New("AI_PROMPT cannot be empty")
	}

	if reflect.TypeOf(input).Kind() != reflect.String {
		return "", errors.New("AI_PROMPT must be a string")
	}

	maxLength := 120

	if len(input) > maxLength {
		return "", errors.New(
			"AI_PROMPT must have less than 120 characters",
		)
	}

	return input, nil
}

func (v *Validator) ModelTemperature(input string) (string, error) {
	if input == "" {
		return "", errors.New("AI_TEMPERATURE cannot be empty")
	}

	if reflect.TypeOf(input).Kind() != reflect.String {
		return "", errors.New("AI_TEMPERATURE must be a string")
	}

	bits := 64
	temp, err := strconv.ParseFloat(input, bits)
	if err != nil {
		log.Fatal(err)
	}

	if temp < 0.1 || temp > 1 || temp*10 != float64(int(temp*10)) {
		return "", errors.New(
			"AI_TEMPERATURE must be a multiple of 0.1, and be between 0.1 and 1.0",
		)
	}

	return input, nil
}

func (v *Validator) ModelBaseURL(input string) (string, error) {
	local := "http://localhost:"

	if input != local {
		return "", errors.New("AI_BASE_URL must be \"http://localhost:\"")
	}

	return input, nil
}
