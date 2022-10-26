package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type TaskReq struct {
	ReqType      string `json:"reqType"`
	UserId       string `json:"userId"`
	BusinessLine string `json:"businessLine"`
}

type WorkFlowConfigRes struct {
	Name        string `json:"name"`
	TargetQueue string `json:"targetQueue"`
	Priority    string `json:"priority"`
}

type TaskQueueReq struct {
	TaskType  string `json:"taskType"`
	UserId    string `json:"userId"`
	QueueName string `json:"queueName"`
	Priority  string `json:"priority"`
}

type TaskQueueResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

var tracer = otel.Tracer("gin-server")

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(otelgin.Middleware("router-service"))
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	// Ping test
	r.GET("/health", func(c *gin.Context) {

		sugar.Infow("inside health func")

		c.String(http.StatusOK, "Ping test successful!")
	})

	r.POST("/route", func(c *gin.Context) {

		logger.Info("inside route handler")

		var taskReq TaskReq

		if err := c.BindJSON(&taskReq); err != nil {
			return
		}

		_, span := tracer.Start(c.Request.Context(), "getUser", oteltrace.WithAttributes(attribute.String("id", "arul-test")))
		defer span.End()

		logger.Info("Received Task Request", zap.Any("requests", taskReq))

		logger.Info("Invoking Workflow engine to retrive config")
		//response, err := http.Get("http://workflow-engine:8080/workflow/config")

		response, err := SendRequest(c, "GET", "http://workflow-engine:8080/workflow/config", nil)

		//response, err := SendRequest(c, "GET", "http://localhost:8082/workflow/config", nil)

		if err != nil {
			logger.Error(err.Error())
			return
		}

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			logger.Fatal(err.Error())
		}

		var workFlowConfigResponse WorkFlowConfigRes
		json.Unmarshal(responseData, &workFlowConfigResponse)

		logger.Info("WorkFlow Engine Config response", zap.Any("workFlowConfig", workFlowConfigResponse))

		taskQueueReq := TaskQueueReq{
			TaskType:  taskReq.ReqType,
			UserId:    taskReq.UserId,
			QueueName: workFlowConfigResponse.TargetQueue,
			Priority:  workFlowConfigResponse.Priority,
		}

		postBody, _ := json.Marshal(taskQueueReq)

		//reqBody := bytes.NewBuffer(postBody)

		//logger.Info("Invoking Task Queue for scheduling", zap.Any("reqBody", reqBody))

		//taskQueueHTTPResponse, err := http.Post("http://task-queue:80/queue", "application/json", reqBody)

		taskQueueHTTPResponse, err := SendRequest(c, "POST", "http://task-queue:80/queue", postBody)

		//taskQueueHTTPResponse, err := SendRequest(c, "POST", "http://localhost:8083/queue", postBody)

		if err != nil {
			logger.Fatal(err.Error())
		}

		taskSCheduleResponseData, err := ioutil.ReadAll(taskQueueHTTPResponse.Body)

		var taskQueueResponse TaskQueueResponse
		json.Unmarshal(taskSCheduleResponseData, &taskQueueResponse)

		logger.Info("route final response", zap.Any("response", taskQueueResponse))

		c.JSON(http.StatusOK, taskQueueResponse)
	})

	return r
}

func main() {
	cleanup := initTracer()
	defer cleanup(context.Background())

	r := setupRouter()
	r.Run(":8080")
}
