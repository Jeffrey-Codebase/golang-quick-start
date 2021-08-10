package main

import (
	"log"
	"os"

	"github.com/Jeffrey-Codebase/hrbrain-go-assignment/config"
	"github.com/Jeffrey-Codebase/hrbrain-go-assignment/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("GITHUB_TOKEN") == "" {
		log.Fatalln("Please store the github access token in env GITHUB_TOKEN")
	}
	config := config.GetConfig()
	engine := gin.Default()
	engine = routes.GetRepoRoute(engine)
	engine.Run(":" + config.Port)
}
