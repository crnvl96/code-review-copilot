package main

import (
	"github.com/omega-energia/code-review-copilot/internal/contexts"
	"github.com/omega-energia/code-review-copilot/internal/model"
	"github.com/omega-energia/code-review-copilot/pkg/env"
)

func main() {
	e := env.Retrieve()

	ctx := contexts.ModelContext()
	llm := model.Spec()
	prompt := e.AiPrompt + "\n"
	temp := model.Temperature()
	str := model.SteamingFunc

	model.Generate(ctx, llm, prompt, temp, str)
}
