package services

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Jeffrey-Codebase/hrbrain-go-assignment/config"
	customErrors "github.com/Jeffrey-Codebase/hrbrain-go-assignment/internal/errors"
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
	service := NewGithubService(getDefaultGithubClient())
	result, err := service.GetGithubRepo(user, repo)
	assert.Nil(t, err)
	assert.True(t, result.StarCount > 0)
	assert.True(t, len(result.Follower) == result.StarCount)
	assert.Equal(t, nil, err)
}

func TestGetRepoWithBadUser(t *testing.T) {
	service := NewGithubService(getDefaultGithubClient())
	_, err := service.GetGithubRepo("baduser", repo)
	assert.NotNil(t, err)
}

func TestGetRepoWithBadRepo(t *testing.T) {
	service := NewGithubService(getDefaultGithubClient())
	_, err := service.GetGithubRepo(user, "badrepo")
	assert.NotNil(t, err)
}

func TestGetRepoTimeoutAndRetry(t *testing.T) {
	service := NewGithubService(getGithubClient(time.Microsecond))
	_, err := service.GetGithubRepo(user, repo)
	assert.NotNil(t, err)
}

func TestGetRepoRateLimitError(t *testing.T) {
	service := NewGithubService(getDefaultGithubClient())
	service.rateLimitResetTime = time.Now().Add(time.Hour)
	_, err := service.GetGithubRepo(user, repo)
	var rateLimitError *customErrors.RateLimitError
	assert.NotNil(t, err)
	assert.True(t, errors.As(err, &rateLimitError))

}

func getDefaultGithubClient() *github.Client {
	if githubClient != nil {
		return githubClient
	}
	githubClient = getGithubClient(0)
	return githubClient
}

func getGithubClient(timeoutSpec time.Duration) *github.Client {

	if os.Getenv("GITHUB_TOKEN") == "" {
		log.Fatalln("Please store the github access token in env GITHUB_TOKEN")
	}
	token := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	if timeoutSpec == 0 {
		config := config.GetConfig()
		timeoutSpec = time.Duration(config.TimeoutMS) * time.Millisecond
	}
	httpClient := &http.Client{
		Transport: oauth2.NewClient(ctx, ts).Transport,
		Timeout:   timeoutSpec,
	}
	return github.NewClient(httpClient)
}
