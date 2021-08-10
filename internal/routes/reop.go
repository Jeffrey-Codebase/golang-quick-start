package routes

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Jeffrey-Codebase/hrbrain-go-assignment/config"
	"github.com/Jeffrey-Codebase/hrbrain-go-assignment/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func GetRepoRoute(r *gin.Engine) *gin.Engine {
	githubService := getGithubService()

	r.GET("/repos", func(c *gin.Context) {
		user := c.Query("user")
		repo := c.Query("repo")
		if user == "" || repo == "" {
			c.String(http.StatusBadRequest, "Bad Request")
			return
		}
		result, err := githubService.GetGithubRepo(user, repo)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, result)
	})

	return r
}

func getGithubService() *services.GithubService {
	config := config.GetConfig()

	token := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := &http.Client{
		Transport: oauth2.NewClient(ctx, ts).Transport,
		Timeout:   time.Duration(config.TimeoutMS) * time.Millisecond,
	}

	client := github.NewClient(httpClient)
	return services.NewGithubService(client)
}
