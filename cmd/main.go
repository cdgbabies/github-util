package main

import (
	"context"
	"os"

	"github.com/google/go-github/v49/github"
	"golang.org/x/oauth2"
)

type GitHubWorkflowClient interface {
	CreateWorkflowDispatchEventByFileName(ctx context.Context, owner, repo, workflowFileName string, event github.CreateWorkflowDispatchEventRequest) (*github.Response, error)
}

func getGitHubClient(ctx context.Context) *github.ActionsService {
	token := os.Getenv("TOKEN")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc).Actions
}

func main() {
	client := getGitHubClient(context.Background())
	ownerName := "cdgbabies"
	workflowRequest := github.CreateWorkflowDispatchEventRequest{
		Ref: "main",
	}
	workflowFileName := "build-upload-to-s3.yml"
	repositoryNames := []string{"list-blogs-lambda", "testimonials-ddb-update-handler-lambda", "blogs-ddb-update-lambda", "blogs-upload-handler-lambda", "add-testimonial-lambda"}

	for _, repoName := range repositoryNames {
		_, err := client.CreateWorkflowDispatchEventByFileName(context.TODO(), ownerName, repoName, workflowFileName, workflowRequest)
		if err != nil {
			panic(err)
		}
	}

}
