package main

import (
	"github.com/omega-energia/code-review-copilot/internal/contexts"
	"github.com/omega-energia/code-review-copilot/internal/model"
	"github.com/omega-energia/code-review-copilot/pkg/spec"
	"github.com/omega-energia/code-review-copilot/pkg/validation"
)

func main() {
	validator := validation.NewValidator()
	spec := spec.NewSpec(validator)
	ctx := contexts.NewModelContext()
	aiModel := model.NewModel(spec)

	envVars := spec.FromEnv()

	aiModel.GetResponse(
		ctx.Start(),
		aiModel.GenerateSpec(),
		envVars.AiPrompt+"\n",
		aiModel.GenerateTemperature(),
		aiModel.StreamingFunc,
	)
}
