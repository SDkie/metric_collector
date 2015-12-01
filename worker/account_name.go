package worker

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/SDkie/metric_collector/logger"
	"github.com/SDkie/metric_collector/model"
)

type WorkerAccountName struct {
	name string
}

func (h WorkerAccountName) Run() {
	logger.Infof("[%s] worker started at %s", h.name, time.Now().UTC())

	ch, err := RabbitChannel.Consume(
		Q_ACCOUNT_NAME, // queue
		"",             // consumer
		false,          // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		logger.Errf("[%s] Error while creating Consume channel, %s", h.name, err)
		return
	}

	for {
		select {

		case msg := <-ch:
			logger.Infof("[%s] Received metric: %s", h.name, msg.Body)
			metricData := new(model.MetricPg)

			err = json.Unmarshal(msg.Body, metricData)
			if err != nil {
				logger.Errf("[%s] Error while doing JSON Unmarshal, %s", h.name, err)
				msg.Reject(true)
				break
			}

			err = metricData.Insert()
			if err == nil {
				logger.Infof("[%s] Metric successfully added to psql - %s", h.name, msg.Body)
				msg.Ack(false)
			} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				logger.Infof("[%s] %s already exists in psql", h.name, metricData.Username)
				msg.Ack(false)
			} else {
				logger.Errf("[%s] Error while Inserting data into psql, %s", h.name, err)
				msg.Reject(true)
			}
		case <-time.After(1 * time.Minute):
			logger.Info("[%s] Worker completed at %s", h.name, time.Now().UTC())
			return
		}
	}
}
