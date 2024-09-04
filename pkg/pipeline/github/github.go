package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/crnvl96/code-review-copilot/pkg/config"
	"github.com/crnvl96/code-review-copilot/pkg/tinyllama"
	"github.com/google/go-github/v64/github"
	"golang.org/x/oauth2"
)

const (
	accessToken = "ACTION_ACCESS_TOKEN"
	repoOwner   = "REPOSITORY_OWNER"
	repoName    = "REPOSITORY_NAME"
	prNumber    = "PULL_REQUEST_NUMBER"
)

func Generate() error {
	tk := os.Getenv(accessToken)
	if tk == "" {
		err := fmt.Sprintf(
			"%s environment variable must be set in repository settings",
			accessToken,
		)
		return errors.New(err)
	}

	owner := os.Getenv(repoOwner)
	repo := os.Getenv(repoName)

	prNumber := os.Getenv(prNumber)
	prNumberInt, err := strconv.Atoi(prNumber)
	if err != nil {
		return err
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: tk})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	files, _, err := client.PullRequests.ListFiles(ctx, owner, repo, prNumberInt, nil)
	if err != nil {
		return err
	}

	config, err := config.GenerateConfig()
	if err != nil {
		return err
	}

	tinyLlama, err := tinyllama.GenerateModel(config)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileContent, err := os.ReadFile(*file.Filename)
		if err != nil {
			log.Printf("Failed to read file %s: %v", *file.Filename, err)
			continue
		}

		data := fmt.Sprintf("```typescript\n%v\n```", fileContent)

		res, err := tinyLlama(data)
		if err != nil {
			return err
		}

		commentBody := fmt.Sprintf("Review for %s:\n\n%s", *file.Filename, res)

		comment := &github.IssueComment{
			Body: github.String(commentBody),
		}

		_, _, err = client.Issues.CreateComment(ctx, owner, repo, prNumberInt, comment)
		if err != nil {
			return err
		}
	}

	return nil
}
