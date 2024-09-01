package github

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
)

type fileContent struct {
	filename string
	content  string
}

func processFileContents(
	ctx context.Context,
	client *github.Client,
	owner, repo string,
	prNumber int,
	fileContents []fileContent,
) {
	commentBody := "Files in the PR:\n"

	for _, fileContent := range fileContents {
		commentBody += fmt.Sprintf(
			"\n## %s\n```\n%s\n```\n",
			fileContent.filename,
			fileContent.content,
		)
	}

	comment := &github.IssueComment{
		Body: &commentBody,
	}

	_, _, err := client.Issues.CreateComment(ctx, owner, repo, prNumber, comment)
	if err != nil {
		log.Fatalf("Error creating comment: %v", err)
	}
}
