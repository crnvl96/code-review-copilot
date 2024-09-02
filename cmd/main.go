package main

import (
	"log"

	"github.com/omega-energia/code-review-copilot/pkg/model/tinyllama"
)

func main() {
	tinyllama := tinyllama.NewTinyLlamaInstance()

	err := tinyllama.Run()
	if err != nil {
		log.Fatal(err)
	}
}
