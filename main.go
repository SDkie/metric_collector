package main

import (
	"os"

	"github.com/SDkie/metric_collector/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()
	logger.Info("Logger Successfully Initialize")
	gin.SetMode(os.Getenv("MODE"))

	router := getRouter()
	router.Run(":" + os.Getenv("PORT"))
}
