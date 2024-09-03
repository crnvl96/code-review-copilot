package github

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/go-github/v64/github"
	"golang.org/x/oauth2"
)

func main() {
	// Get the GITHUB_TOKEN from environment variables
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN environment variable not set")
	}

	// Get the repository and pull request information from environment variables
	owner := os.Getenv("GITHUB_REPOSITORY_OWNER")
	repo := os.Getenv("GITHUB_REPOSITORY_NAME")
	prNumber := os.Getenv("GITHUB_PULL_REQUEST_NUMBER")

	// Convert prNumber to an integer
	prNumberInt, err := strconv.Atoi(prNumber)
	if err != nil {
		log.Fatalf("Failed to convert PR number to int: %v", err)
	}

	// Set up the GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// List files in the pull request
	files, _, err := client.PullRequests.ListFiles(ctx, owner, repo, prNumberInt, nil)
	if err != nil {
		log.Fatalf("Failed to list files in the pull request: %v", err)
	}

	// Iterate over the files and create a comment for each file
	for _, file := range files {
		comment := &github.IssueComment{
			Body: github.String(fmt.Sprintf("File: %s", *file.Filename)),
		}

		_, _, err := client.Issues.CreateComment(ctx, owner, repo, prNumberInt, comment)
		if err != nil {
			log.Fatalf("Failed to create comment for file %s: %v", *file.Filename, err)
		}

		log.Printf("Comment created for file: %s", *file.Filename)
	}
}
