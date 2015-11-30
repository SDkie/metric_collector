package worker

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/SDkie/metric_collector/logger"
	"github.com/SDkie/metric_collector/model"
)

type WorkerAccountName struct {
}

func (h WorkerAccountName) Run() {
	logger.Info("AccountName worker started at ", time.Now().UTC())

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
		logger.Errf("Error while creating Consume channel in AccountName worker , %s", err)
		return
	}

	for {
		select {

		case msg := <-ch:
			logger.Infof("Received metric: %s", msg.Body)
			metricData := new(model.MetricPg)

			err = json.Unmarshal(msg.Body, metricData)
			if err != nil {
				logger.Errf("Error while doing JSON Unmarshal, %s", err)
				msg.Reject(true)
				break
			}

			err = metricData.Insert()
			if err == nil {
				logger.Info("Metric successfully added to psql")
				msg.Ack(false)
			} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				logger.Infof("Account Name, %s already exists in DB", metricData.Username)
				msg.Ack(false)
			} else {
				logger.Errf("Error while Inserting data into psql, %s", err)
				msg.Reject(true)
			}
		case <-time.After(1 * time.Minute):
			logger.Info("HourlyLog worker completed at ", time.Now().UTC())
			return
		}
	}
}
