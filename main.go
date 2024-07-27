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
	AI_CPU_PORT       = "AI_CPU_PORT"
	AI_GPU_PORT       = "AI_GPU_PORT"
	AI_PROMPT         = "AI_PROMPT"
	AI_TEMPERATURE    = "AI_TEMPERATURE"
	AI_PRINT_BY_CHUNK = "AI_PRINT_BY_CHUNK"
	AI_RUN_ON_GPU     = "AI_RUN_ON_GPU"
)

func print() {
	fmt.Print("test")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	shouldRunOnGPU, err := strconv.ParseBool(os.Getenv(AI_RUN_ON_GPU))
	if err != nil {
		log.Fatal(err)
	}

	serverURL := func() ollama.Option {
		baseURL := "http://localhost:"
		var URL string

		if shouldRunOnGPU {
			URL = baseURL + os.Getenv(AI_GPU_PORT)

			return ollama.WithServerURL(URL) // 34%/66% CPU/GPU based
		}

		URL = baseURL + os.Getenv(AI_CPU_PORT)

		return ollama.WithServerURL(URL) // 100% CPU based
	}

	llm, err := ollama.New(
		ollama.WithModel(os.Getenv(AI_MODEL_NAME)),
		serverURL(),
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
