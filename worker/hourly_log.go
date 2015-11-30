package worker

import (
	"encoding/json"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/SDkie/metric_collector/logger"
	"github.com/SDkie/metric_collector/model"
)

type WorkerHourlyLog struct {
}

func (h WorkerHourlyLog) Run() {
	logger.Info("HourlyLog worker started at ", time.Now().UTC())

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
		logger.Errf("Error while creating Consume channel in HourlyLog worker , %s", err)
		return
	}

	for {
		select {

		case msg := <-ch:
			logger.Infof("Received metric: %s", msg.Body)
			metricData := new(model.MetricMongo)
			metricData.Id = bson.NewObjectId()

			err = json.Unmarshal(msg.Body, metricData)
			if err != nil {
				logger.Errf("Error while doing JSON Unmarshal, %s", err)
				msg.Reject(true)
				break
			}

			err = metricData.Insert()
			if err != nil {
				logger.Errf("Error while Inserting data into Mongo, %s", err)
				msg.Reject(true)
				break
			}
			msg.Ack(false)
			logger.Info("Metric successfully added to mongo db")

		case <-time.After(1 * time.Minute):
			logger.Info("HourlyLog worker completed at ", time.Now().UTC())
			return
		}
	}
}
