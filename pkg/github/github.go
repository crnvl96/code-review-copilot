package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/go-github/v64/github"
	"github.com/omega-energia/code-review-copilot/pkg/tinyllama"
	"golang.org/x/oauth2"
)

const (
	AccessToken = "ACESS_TOKEN"
	RepoOwner   = "REPOSITORY_OWNER"
	RepoName    = "REPOSITORY_NAME"
	PrNumber    = "PULL_REQUEST_NUMBER"
)

func Generate() error {
	// This variable is retrieved from repository secrets, instead of a `.env` file
	token := os.Getenv(AccessToken)
	if token == "" {
		err := fmt.Sprintf(
			"%s environment variable must be set in respository settings",
			AccessToken,
		)
		return errors.New(err)
	}

	owner := os.Getenv(RepoOwner)
	repo := os.Getenv(RepoName)

	prNumberInt, err := strconv.Atoi(os.Getenv(PrNumber))
	if err != nil {
		return errors.New("Failed to convert PR number to int")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	files, _, err := client.PullRequests.ListFiles(ctx, owner, repo, prNumberInt, nil)
	if err != nil {
		return errors.New("Failed to list files in the pull request")
	}

	for _, file := range files {
		fileContent, err := os.ReadFile(*file.Filename)
		if err != nil {
			log.Printf("Failed to read file %s: %v", *file.Filename, err)
			continue
		}

		data := fmt.Sprintf("```typescript\n%s\n```", fileContent)

		fmt.Println("I got here!!!")

		res, err := tinyllama.Run(data)
		if err != nil {
			return err
		}

		commentBody := fmt.Sprintf("Review for %s:\n%s", *file.Filename, res)

		comment := &github.IssueComment{
			Body: github.String(commentBody),
		}

		_, _, err = client.Issues.CreateComment(ctx, owner, repo, prNumberInt, comment)
		if err != nil {
			return errors.New("Failed to create comment")
		}
	}

	return nil
}
