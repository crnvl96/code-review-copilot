package main

import (
	"github.com/omega-energia/code-review-copilot/internal/contexts"
	"github.com/omega-energia/code-review-copilot/internal/model"
	"github.com/omega-energia/code-review-copilot/pkg/env"
)

func main() {
	e := env.Retrieve()

	model.Generate(contexts.ModelContext(), model.Spec(), e.AiPrompt+"\n", model.Temperature(), model.StreamingFunc)
}
