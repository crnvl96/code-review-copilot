package main

import (
	"log"

	"github.com/crnvl96/code-review-copilot/pkg/codereview"
)

func main() {
	error := codereview.Generate()
	if error != nil {
		log.Fatal(error)
	}
}
