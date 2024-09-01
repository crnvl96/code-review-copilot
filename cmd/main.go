package main

import "github.com/omega-energia/code-review-copilot/pkg/github"

func main() {
	// tinyllama := tinyllama.NewTinyLlamaInstance()
	//
	// err := tinyllama.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//	loader := github.NewGithubSecretsLoader()

	loader := github.NewGithubSecretsLoader()
	validator := github.NewGithubSecretsValidator(loader)
	parser := github.NewGithubSecretsParser(validator)
	configurator := github.NewGithubSecretsConfigurator(parser)
	starter := github.NewGithubClientStarter(configurator)

	starter.Start()
}
