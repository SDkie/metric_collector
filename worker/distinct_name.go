package worker

import (
	"encoding/json"
	"time"

	"github.com/SDkie/metric_collector/logger"
	"github.com/SDkie/metric_collector/model"
)

type WorkerDistinctName struct {
	name string
}

func (d WorkerDistinctName) Run() {
	logger.Infof("[%s] worker started at %s", d.name, time.Now().UTC())

	ch, err := RabbitChannel.Consume(
		Q_DISTINCT_NAME, // queue
		"",              // consumer
		false,           // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		logger.Errf("[%s] Error while creating Consume channel, %s", d.name, err)
		return
	}

	for {
		select {

		case msg := <-ch:
			logger.Infof("[%s] Received metric: %s", d.name, msg.Body)
			metricData := new(model.MetricRedis)

			err = json.Unmarshal(msg.Body, metricData)
			if err != nil {
				logger.Errf("[%s] Error while doing JSON Unmarshal, %s", d.name, err)
				msg.Reject(true)
				break
			}

			err = metricData.Insert()
			if err != nil {
				logger.Errf("[%s] Error while Inserting data into Redis, %s", d.name, err)
				msg.Reject(true)
				break
			}
			msg.Ack(false)
			logger.Infof("[%s] Metric successfully added to Redis - %s", d.name, msg.Body)

		case <-time.After(1 * time.Minute):
			logger.Infof("[%s] Worker completed at %s", d.name, time.Now().UTC())
			return
		}
	}
}
