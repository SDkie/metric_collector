package main

import (
	"github.com/SDkie/metric_collector/controller"
	"github.com/SDkie/metric_collector/http_metric"
	"github.com/SDkie/metric_collector/worker"
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", controller.Root)
	router.POST("/metric", http_metric.HttpMetricMiddleware(), controller.PostMetric)
	router.LoadHTMLGlob("../../../github.com/bamzi/jobrunner/views/Status.html")
	router.GET("/jobrunner/html", worker.JobHtml)
	return router
}
