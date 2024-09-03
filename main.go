package main

import (
	"fmt"
	"log"

	gh "github.com/crnvl96/code-review-copilot/pkg/github"
)

func main() {
	error := gh.Generate()
	if error != nil {
		fmt.Println(error)
		log.Fatal(error)
	}
}
