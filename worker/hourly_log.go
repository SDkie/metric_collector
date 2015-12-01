package worker

import (
	"encoding/json"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/SDkie/metric_collector/logger"
	"github.com/SDkie/metric_collector/model"
)

type WorkerHourlyLog struct {
	name string
}

func (h WorkerHourlyLog) Run() {
	logger.Infof("[%s] worker started at %s", h.name, time.Now().UTC())

	ch, err := RabbitChannel.Consume(
		Q_HOURLY_LOG, // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		logger.Errf("[%s] Error while creating Consume channel in HourlyLog worker, %s", h.name, err)
		return
	}

	for {
		select {

		case msg := <-ch:
			logger.Infof("[%s] Received metric: %s", h.name, msg.Body)
			metricData := new(model.MetricMongo)
			metricData.Id = bson.NewObjectId()

			err = json.Unmarshal(msg.Body, metricData)
			if err != nil {
				logger.Errf("[%s] Error while doing JSON Unmarshal, %s", h.name, err)
				msg.Reject(true)
				break
			}

			err = metricData.Insert()
			if err != nil {
				logger.Errf("[%s] Error while Inserting data into Mongo, %s", h.name, err)
				msg.Reject(true)
				break
			}
			msg.Ack(false)
			logger.Infof("[%s] Metric successfully added to mongoDb - %s", h.name, msg.Body)

		case <-time.After(1 * time.Minute):
			logger.Infof("[%s] Worker completed at %s", h.name, time.Now().UTC())
			return
		}
	}
}
