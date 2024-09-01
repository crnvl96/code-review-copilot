package tinyllama

import (
	"strconv"

	"github.com/omega-energia/code-review-copilot/pkg/model"
)

type tinyLlamaArgsParser struct {
	validator model.Validator
}

func newTinyLlamaArgsParser(validator model.Validator) *tinyLlamaArgsParser {
	return &tinyLlamaArgsParser{
		validator: validator,
	}
}

func (t *tinyLlamaArgsParser) Parse() (model.ModelParams, error) {
	params, err := t.validator.Validate()

	temp, err := strconv.ParseFloat(params.Temp, 64)
	if err != nil {
		return model.ModelParams{}, err
	}

	return model.ModelParams{
		Name:   params.Name,
		Port:   params.Port,
		Prompt: params.Prompt,
		Temp:   temp,
	}, nil
}
