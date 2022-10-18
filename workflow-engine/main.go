package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WorkFlowConfig struct {
	Name            string `json:"name"`
	TargetQueueName string `json:"targetQueue"`
	Priority        string `json:"priority"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// return Workflow information for router
	r.GET("/workflow/config", func(c *gin.Context) {
		logger, _ := zap.NewProduction()
		logger.Info("inside Get WorkFlow Config func")
		var config = WorkFlowConfig{
			Name:            "test-workflow",
			TargetQueueName: "test-queue",
			Priority:        "high",
		}
		logger.Info("Returned Workflow Config Details",
			// Structured context as strongly typed Field values.
			zap.String("Name", config.Name),
			zap.String("TargetQueue", config.TargetQueueName),
			zap.String("Priority", config.Priority),
		)
		c.JSON(http.StatusOK, config)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
