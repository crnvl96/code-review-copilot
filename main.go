package main

import (
	"log"

	gh "github.com/crnvl96/code-review-copilot/pkg/pipeline/github"
)

func main() {
	error := gh.Generate()
	if error != nil {
		log.Fatal(error)
	}
}
