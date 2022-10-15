package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/enqueue", func(c *gin.Context) {

		//get config from work flow engine service

		//enqueue task info..

		c.String(http.StatusOK, "Task Queued Successfully")
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
