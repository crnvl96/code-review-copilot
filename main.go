package main

import (
	"fmt"

	gh "github.com/omega-energia/code-review-copilot/pkg/github"
)

func main() {
	fmt.Println("Starting job execution")
	gh.Generate()
}
