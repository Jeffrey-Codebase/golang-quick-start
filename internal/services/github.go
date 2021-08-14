package services

import (
	"context"
	"errors"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/Jeffrey-Codebase/hrbrain-go-assignment/config"
	customErrors "github.com/Jeffrey-Codebase/hrbrain-go-assignment/internal/errors"
	"github.com/Jeffrey-Codebase/hrbrain-go-assignment/internal/utils"
	"github.com/google/go-github/github"
)

type Response struct {
	StarCount int      `json:"star_count"`
	Follower  []string `json:"followers"`
}

type GithubService struct {
	client             *github.Client
	rateLimitResetTime time.Time
}

func NewGithubService(client *github.Client) *GithubService {
	return &GithubService{client: client}
}

func (gs *GithubService) GetGithubRepo(username string, reponame string) (*Response, error) {
	// check rate limit
	if gs.rateLimitResetTime.After(time.Now()) {
		return nil, customErrors.NewRateLimitError(gs.rateLimitResetTime)
	}

	// get repository from github

	ctx := context.Background()
	repo, resp, err := gs.getRepository(ctx, username, reponame)

	if err != nil || resp.StatusCode != http.StatusOK {
		// if reach github rate limitation
		var githubRateLimitError *github.RateLimitError
		if err != nil && errors.As(err, &githubRateLimitError) {
			gs.rateLimitResetTime = resp.Rate.Reset.Time
		}

		log.Print("Get Repository Failed. ", utils.GetStatusCode(resp), err)
		return nil, errors.New("Get Repository Failed.")
	}
	starCount := repo.StargazersCount

	// get stargazers from github
	follower, err := gs.getStargazers(ctx, username, reponame, *starCount)
	if err != nil {
		var githubRateLimitError *github.RateLimitError
		if errors.As(err, &githubRateLimitError) {
			gs.rateLimitResetTime = resp.Rate.Reset.Time
		}
		return nil, err
	}

	result := &Response{StarCount: *starCount, Follower: follower}
	return result, nil
}

func (gs *GithubService) getStargazers(ctx context.Context, username string, reponame string, starCount int) ([]string, error) {
	channel := make(chan []*github.Stargazer)
	defer close(channel)

	var follower []string
	totalPages := int(math.Ceil(float64(starCount) / 100))
	for page := 1; page <= totalPages; page++ {
		// use goroutine to fetch data parallelly
		go gs.getStargazersByPage(ctx, username, reponame, page, channel)
	}
	success := 0
	for finish := 0; finish < totalPages; finish++ {
		stargazers := <-channel
		if stargazers != nil {
			success++
			for _, stargazer := range stargazers {
				follower = append(follower, *stargazer.User.Login)
			}
		}
	}
	if success != totalPages {
		return nil, errors.New("Get Stargazers Failed.")
	}
	return follower, nil
}

func (gs *GithubService) getStargazersByPage(ctx context.Context, username string, reponame string, pageNo int, ch chan<- []*github.Stargazer) {
	var option github.ListOptions
	option.Page = pageNo
	option.PerPage = 100
	maxAttempt := config.GetConfig().MaxAttempt
	attempt := 1
	stargazers, resp, err := gs.client.Activity.ListStargazers(ctx, username, reponame, &option)
	// retry for timeout error or 500 reture code
	for attempt < maxAttempt && (os.IsTimeout(err) || resp.StatusCode == http.StatusInternalServerError) {
		attempt++
		log.Print("Get Stargazers Failed. Retry attempt ", attempt)
		stargazers, resp, err = gs.client.Activity.ListStargazers(ctx, username, reponame, &option)
	}
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Print("Get Stargazers Failed. ", utils.GetStatusCode(resp), err)
		ch <- nil
		return
	}
	ch <- stargazers
}

func (gs *GithubService) getRepository(ctx context.Context, username string, reponame string) (*github.Repository, *github.Response, error) {
	maxAttempt := config.GetConfig().MaxAttempt
	attempt := 1
	repo, resp, err := gs.client.Repositories.Get(ctx, username, reponame)
	// retry for timeout error or 500 reture code
	for attempt < maxAttempt && (os.IsTimeout(err) || resp.StatusCode == http.StatusInternalServerError) {
		attempt++
		log.Print("Get Repository Failed. Retry attempt ", attempt)
		repo, resp, err = gs.client.Repositories.Get(ctx, username, reponame)
	}
	return repo, resp, err
}
