package main

import (
	"log"

	"github.com/omega-energia/code-review-copilot/ollama/config"
	"github.com/omega-energia/code-review-copilot/ollama/model"
)

func main() {
	var (
		tinyLlamaLoader    = config.NewTinyLlamaLoader()
		tinyLlamaValidator = config.NewTinyLlamaValidator()
		tinyLlamaParser    = config.NewTinyLlamaParser()
	)

	args, err := tinyLlamaLoader.Load()
	if err != nil {
		log.Fatal(err)
	}

	valid, err := tinyLlamaValidator.Validate(args)
	if err != nil {
		log.Fatal(err)
	}

	config, err := tinyLlamaParser.Parse(valid)
	if err != nil {
		log.Fatal(err)
	}

	ollama := model.NewTinyLlama(config)

	err = ollama.Run()
	if err != nil {
		log.Fatal(err)
	}
}
