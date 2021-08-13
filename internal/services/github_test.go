package services

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Jeffrey-Codebase/hrbrain-go-assignment/config"
	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

var (
	user = "codecov"
	repo = "example-go"
)

var githubClient *github.Client

func TestGetRepoSuccess(t *testing.T) {
	service := NewGithubService(getGithubClient())
	result, err := service.GetGithubRepo(user, repo)
	assert.Nil(t, err)
	assert.True(t, result.StarCount > 0)
	assert.True(t, len(result.Follower) == result.StarCount)
	assert.Equal(t, nil, err)
}

func TestGetRepoWithBadUser(t *testing.T) {
	service := NewGithubService(getGithubClient())
	_, err := service.GetGithubRepo("baduser", repo)
	assert.NotNil(t, err)
}

func TestGetRepoWithBadRepo(t *testing.T) {
	service := NewGithubService(getGithubClient())
	_, err := service.GetGithubRepo(user, "badrepo")
	assert.NotNil(t, err)
}

func getGithubClient() *github.Client {

	if githubClient != nil {
		return githubClient
	}
	if os.Getenv("GITHUB_TOKEN") == "" {
		log.Fatalln("Please store the github access token in env GITHUB_TOKEN")
	}
	token := os.Getenv("GITHUB_TOKEN")
	config := config.GetConfig()
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := &http.Client{
		Transport: oauth2.NewClient(ctx, ts).Transport,
		Timeout:   time.Duration(config.TimeoutMS) * time.Millisecond,
	}
	githubClient = github.NewClient(httpClient)
	return githubClient
}
