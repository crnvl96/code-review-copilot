package main

import (
	"log"

	"github.com/omega-energia/code-review-copilot/pkg/model"
	"github.com/omega-energia/code-review-copilot/pkg/model/tinyllama"
)

func main() {
	tinyLlama := tinyllama.NewTinyLlama()
	myModel := model.NewModel(tinyLlama)

	err := myModel.Instance.Run()
	if err != nil {
		log.Fatal(err)
	}
}
