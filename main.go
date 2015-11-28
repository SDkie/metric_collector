package main

import (
	"os"

	"github.com/SDkie/metric_collector/logger"
	"github.com/SDkie/metric_collector/worker"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()
	worker.InitRabbitMQ()
	gin.SetMode(os.Getenv("MODE"))

	router := getRouter()
	router.Run(":" + os.Getenv("PORT"))
}
