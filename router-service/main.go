package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	// Ping test
	r.GET("/route", func(c *gin.Context) {

		sugar.Infow("inside setupRouter func",
			// Structured context as loosely typed key-value pairs.
			"url", "/route",
			"attempt", 3,
		)

		//get config from work flow engine service

		//enqueue task info.
		c.String(http.StatusOK, "Task Queued Successfully")
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
