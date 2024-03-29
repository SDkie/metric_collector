package main

import (
	"os"

	"github.com/SDkie/metric_collector/http_metric"
	"github.com/SDkie/metric_collector/logger"
	"github.com/SDkie/metric_collector/model"
	"github.com/SDkie/metric_collector/worker"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()
	model.Init()
	worker.InitRabbitMQ()
	worker.InitCron()
	http_metric.Init()
	gin.SetMode(os.Getenv("GIN_MODE"))

	router := getRouter()
	router.Run(":" + os.Getenv("PORT"))
}
