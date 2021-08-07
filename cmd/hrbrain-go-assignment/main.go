package main

import (
	"net/http"

	"github.com/Jeffrey-Codebase/hrbrain-go-assignment/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	githubService := services.GetGithubService()

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, githubService.SayHello())
	})
	router.Run("localhost:8080")
}
