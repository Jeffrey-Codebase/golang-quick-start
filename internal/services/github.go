package services

import (
	"context"
	"errors"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/Jeffrey-Codebase/hrbrain-go-assignment/config"
	"github.com/Jeffrey-Codebase/hrbrain-go-assignment/internal/utils"
	"github.com/google/go-github/github"
)

type Response struct {
	StarCount *int      `json:"star_count"`
	Follower  []*string `json:"followers"`
}

type GithubService struct {
	client *github.Client
}

func NewGithubService(client *github.Client) *GithubService {
	return &GithubService{client: client}
}

func (gs *GithubService) GetGithubRepo(username string, reponame string) (*Response, error) {
	ctx := context.Background()
	// get repository from github
	repo, resp, err := gs.getRepository(ctx, username, reponame)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Print("Get Repository Failed. ", utils.GetStatusCode(resp), err)
		return nil, errors.New("Get Repository Failed.")
	}
	starCount := repo.StargazersCount

	// get stargazers from github
	var follower []*string
	totalPages := int(math.Ceil(float64(*starCount) / 100))
	for page := 1; page <= totalPages; page++ {
		stargazers, resp, err := gs.getStargazersByPage(ctx, username, reponame, page)
		if err != nil || resp.StatusCode != http.StatusOK {
			log.Print("Get Stargazers Failed. ", utils.GetStatusCode(resp), err)
			return nil, errors.New("Get Stargazers Failed.")
		}

		for _, stargazer := range stargazers {
			follower = append(follower, stargazer.User.Login)
		}
	}

	result := &Response{StarCount: starCount, Follower: follower}

	return result, nil
}

func (gs *GithubService) getStargazersByPage(ctx context.Context, username string, reponame string, pageNo int) ([]*github.Stargazer, *github.Response, error) {
	var option github.ListOptions
	option.Page = pageNo
	option.PerPage = 100
	maxAttempt := config.GetConfig().MaxAttempt
	attempt := 1
	stargazers, resp, err := gs.client.Activity.ListStargazers(ctx, username, reponame, &option)
	for attempt < maxAttempt && (os.IsTimeout(err) || resp.StatusCode == http.StatusInternalServerError) {
		attempt++
		log.Print("Get 500 Failed. Retry attempt ", attempt)
		stargazers, resp, err = gs.client.Activity.ListStargazers(ctx, username, reponame, &option)
	}
	return stargazers, resp, err
}

func (gs *GithubService) getRepository(ctx context.Context, username string, reponame string) (*github.Repository, *github.Response, error) {
	maxAttempt := config.GetConfig().MaxAttempt
	attempt := 1
	repo, resp, err := gs.client.Repositories.Get(ctx, username, reponame)
	for attempt < maxAttempt && (os.IsTimeout(err) || resp.StatusCode == http.StatusInternalServerError) {
		attempt++
		log.Print("Get 500 Failed. Retry attempt ", attempt)
		repo, resp, err = gs.client.Repositories.Get(ctx, username, reponame)
	}
	return repo, resp, err
}
