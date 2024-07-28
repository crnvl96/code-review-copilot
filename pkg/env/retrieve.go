package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	AiModelName    string
	AiPort         string
	AiPrompt       string
	AiTemperature  string
	AiPrintByChunk string
	AiBaseUrl      string
}

var (
	aiModelName    = "AI_MODEL_NAME"
	aiPort         = "AI_PORT"
	aiPrompt       = "AI_PROMPT"
	aiTemperature  = "AI_TEMPERATURE"
	aiPrintByChunk = "AI_PRINT_BY_CHUNK"
	aiBaseURL      = "AI_BASE_URL"
)

func Retrieve() Environment {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return Environment{
		AiModelName:    os.Getenv(aiModelName),
		AiPort:         os.Getenv(aiPort),
		AiPrompt:       os.Getenv(aiPrompt),
		AiTemperature:  os.Getenv(aiTemperature),
		AiPrintByChunk: os.Getenv(aiPrintByChunk),
		AiBaseUrl:      os.Getenv(aiBaseURL),
	}
}
