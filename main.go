package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

var (
	AI_MODEL_NAME     = "AI_MODEL_NAME"
	AI_PORT           = "AI_PORT"
	AI_PROMPT         = "AI_PROMPT"
	AI_TEMPERATURE    = "AI_TEMPERATURE"
	AI_PRINT_BY_CHUNK = "AI_PRINT_BY_CHUNK"
)

func print() {
	fmt.Print("test")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	llm, err := ollama.New(
		ollama.WithModel(os.Getenv(AI_MODEL_NAME)),
		ollama.WithServerURL("http://localhost:"+os.Getenv(AI_PORT)),
	)
	if err != nil {
		log.Fatal(err)
	}

	temp, err := strconv.ParseFloat(os.Getenv(AI_TEMPERATURE), 64)
	if err != nil {
		log.Fatal(err)
	}

	shouldPrintByChunk, err := strconv.ParseBool(os.Getenv(AI_PRINT_BY_CHUNK))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		os.Getenv(AI_PROMPT)+"\n",
		llms.WithTemperature(temp),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))

			if shouldPrintByChunk {
				fmt.Print("\n")
			}

			return nil
		}),
	)

	if !shouldPrintByChunk {
		fmt.Print("\n")
	}

	if err != nil {
		log.Fatal(err)
	}

	_ = completion
}
