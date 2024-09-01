package github

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	githubToken = "GITHUB_TOKEN"
	owner       = "GITHUB_REPOSITORY_OWNER"
	repo        = "GITHUB_REPOSITORY"
	prNumber    = "PR_NUMBER"
)

type githubSecretsArgs struct {
	token    string
	owner    string
	repo     string
	prNumber string
}

type githubSecretsParams struct {
	token    string
	owner    string
	repo     string
	prNumber int
	prId     string
}

type loader interface {
	load() (githubSecretsArgs, error)
}

type githubSecretsLoader struct{}

func NewGithubSecretsLoader() *githubSecretsLoader {
	return &githubSecretsLoader{}
}

type validator interface {
	validate() (githubSecretsArgs, error)
}

type githubSecretsValidator struct {
	loader loader
}

func NewGithubSecretsValidator(loader loader) *githubSecretsValidator {
	return &githubSecretsValidator{loader: loader}
}

type parser interface {
	parse() (githubSecretsParams, error)
}

type githubSecretsParser struct {
	validator validator
}

func NewGithubSecretsParser(validator validator) *githubSecretsParser {
	return &githubSecretsParser{validator: validator}
}

type configurator interface {
	getParams() (githubSecretsParams, error)
}

type githubSecretsConfigurator struct {
	parser parser
}

func NewGithubSecretsConfigurator(parser parser) *githubSecretsConfigurator {
	return &githubSecretsConfigurator{parser: parser}
}

func (g *githubSecretsLoader) load() (githubSecretsArgs, error) {
	err := godotenv.Load()
	if err != nil {
		return githubSecretsArgs{}, err
	}

	return githubSecretsArgs{
		token:    os.Getenv(githubToken),
		owner:    os.Getenv(owner),
		repo:     os.Getenv(repo),
		prNumber: os.Getenv(prNumber),
	}, nil
}

func (g *githubSecretsValidator) validate() (githubSecretsArgs, error) {
	args, err := g.loader.load()
	if err != nil {
		return githubSecretsArgs{}, err
	}

	if args.token == "" || args.owner == "" || args.repo == "" || args.prNumber == "" {
		return githubSecretsArgs{}, errors.New(
			"Environment variables GITHUB_TOKEN, GITHUB_REPOSITORY_OWNER, GITHUB_REPOSITORY, and PR_NUMBER must be set",
		)
	}

	return args, nil
}

func (g *githubSecretsParser) parse() (githubSecretsParams, error) {
	args, err := g.validator.validate()
	if err != nil {
		return githubSecretsParams{}, err
	}

	prNum, err := strconv.Atoi(args.prNumber)
	if err != nil {
		return githubSecretsParams{}, errors.New(fmt.Sprintf("Invalid PR number: %v", err))
	}

	return githubSecretsParams{
		token:    args.token,
		repo:     args.repo,
		owner:    args.owner,
		prNumber: prNum,
		prId:     args.prNumber,
	}, nil
}

func (g *githubSecretsConfigurator) getParams() (githubSecretsParams, error) {
	params, err := g.parser.parse()
	if err != nil {
		return githubSecretsParams{}, err
	}

	return params, nil
}
