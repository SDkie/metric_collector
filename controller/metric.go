package controller

import (
	"encoding/json"
	"fmt"
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
	err = addToQueue(metric, worker.Q_ACCOUNT_NAME)
	if err != nil {
		logger.Err(err)
		c.Error(err)
	}

	err = addToQueue(metric, worker.Q_DISTINCT_NAME)
	if err != nil {
		logger.Err(err)
		c.Error(err)
	}

	err = addToQueue(metric, worker.Q_HOURLY_LOG)
	if err != nil {
		logger.Err(err)
		c.Error(err)
	}

	c.JSON(http.StatusOK, metric)
}

func addToQueue(data *model.Metric, queueName string) error {
	dataBytes, _ := json.Marshal(data)

	err := worker.RabbitChannel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     //immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         dataBytes,
		})

	if err != nil {
		return fmt.Errorf("Error while adding task to Queue %s, %s", queueName, err)
	} else {
		return fmt.Errorf("Added task to Queue %s", queueName)
	}

	return err
}
