package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SDkie/metric_collector/logger"
	"github.com/SDkie/metric_collector/model"
	"github.com/SDkie/metric_collector/worker"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
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

	metric.CreatedAt = time.Now().UTC()
	dataBytes, _ := json.Marshal(metric)

	err = worker.RabbitChannel.Publish(
		worker.E_METRIC_EXCHANGE, // exchange
		"",    // routing key
		false, // mandatory
		false, //immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         dataBytes,
		})

	if err != nil {
		logger.Errf("Failed to Publish Message %s, %s", *metric, err)
		c.String(http.StatusInternalServerError, "Failed to Publish Message.")
		return
	}

	logger.Infof("Message is published to exchange, %s", string(dataBytes))
	c.JSON(http.StatusOK, metric)
}
