package service

import (
	"context"

	"github.com/google/go-github/v49/github"
	"golang.org/x/oauth2"
)

type GitHubWorkflowClient interface {
	CreateWorkflowDispatchEventByFileName(ctx context.Context, owner, repo, workflowFileName string, event github.CreateWorkflowDispatchEventRequest) (*github.Response, error)
}

func GetGitHubClient(ctx context.Context) *github.ActionsService {
	token := "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"

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc).Actions
}
