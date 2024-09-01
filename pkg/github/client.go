package github

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/google/go-github/github"
	gh "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type clientStarter interface {
	Start() error
}

type githubClientStarter struct {
	configurator configurator
}

func NewGithubClientStarter(configurator configurator) *githubClientStarter {
	return &githubClientStarter{
		configurator: configurator,
	}
}

func (g *githubClientStarter) Start() error {
	params, err := g.configurator.getParams()
	if err != nil {
		return err
	}

	ctx := context.Background()
	tk := &oauth2.Token{AccessToken: params.token}
	ts := oauth2.StaticTokenSource(tk)
	tc := oauth2.NewClient(ctx, ts)
	client := gh.NewClient(tc)

	files, _, err := client.PullRequests.ListFiles(
		ctx,
		params.owner,
		params.repo,
		params.prNumber,
		nil,
	)
	if err != nil {
		return errors.New(fmt.Sprintf("Error listing PR files: %v", err))
	}

	var fileContents []fileContent

	for _, file := range files {
		content, _, _, err := client.Repositories.GetContents(
			ctx,
			params.owner,
			params.repo,
			*file.Filename,
			&github.RepositoryContentGetOptions{Ref: "refs/pull/" + params.prId + "/head"},
		)
		if err != nil {
			return errors.New(fmt.Sprintf("Error getting file content: %v", err))
		}

		decoded, err := base64.StdEncoding.DecodeString(*content.Content)
		if err != nil {
			return errors.New(fmt.Sprintf("Error decoding file content: %v", err))
		}

		fileContents = append(fileContents, fileContent{
			filename: *file.Filename,
			content:  string(decoded),
		})
	}

	processFileContents(ctx, client, owner, params.repo, params.prNumber, fileContents)

	return nil
}
