package config

import (
	"strconv"
)

type Parser interface {
	Parse(args) (Params, error)
}

type tinyLlamaParser struct{}

func NewTinyLlamaParser() *tinyLlamaParser {
	return &tinyLlamaParser{}
}

func (t *tinyLlamaParser) Parse(a args) (Params, error) {
	temp, err := strconv.ParseFloat(a.temp, 64)
	if err != nil {
		return Params{}, err
	}

	return Params{
		Name:   a.name,
		Port:   a.port,
		Prompt: a.prompt,
		Temp:   temp,
	}, nil
}
