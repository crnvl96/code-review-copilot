package env

import (
	"log"

	"github.com/joho/godotenv"
)

type EnvLoaderInterface interface {
	Load()
}

type EnvLoader struct{}

func NewEnvLoader() *EnvLoader {
	return &EnvLoader{}
}

func (e *EnvLoader) Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
