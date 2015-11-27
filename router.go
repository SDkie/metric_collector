package main

import (
	"github.com/SDkie/metric_collector/controller"
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", controller.Root)
	router.POST("/metric", controller.PostMetric)
	return router
}
