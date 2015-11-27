package controller

import (
	"net/http"

	"github.com/SDkie/metric_collector/logger"
	"github.com/SDkie/metric_collector/model"
	"github.com/gin-gonic/gin"
)

func PostMetric(c *gin.Context) {
	metric := new(model.Metric)
	err := c.BindJSON(metric)
	if err != nil {
		logger.Errf("Error while JSON Binding, %s", err)
		c.String(http.StatusBadRequest, "Invalid JSON, %s", err)
		return
	}

	if metric.Metric == "" || metric.Username == "" || metric.Count == 0 {
		logger.Errf("Invalid metric, %s", *metric)
		c.String(http.StatusBadRequest, "Invalid metric")
		return
	}

	c.JSON(http.StatusOK, metric)
}
