package worker

import (
	"time"

	"github.com/SDkie/metric_collector/logger"
)

type WorkerHourlyLog struct {
}

func (h WorkerHourlyLog) Run() {
	logger.Info("HourlyLog task started at ", time.Now().UTC())

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
			logger.Infof("Received a message: %s", msg.Body)
			msg.Ack(false)
			// TODO: add the data into mongodb
		case <-time.After(1 * time.Minute):
			logger.Info("HourlyLog task completed at ", time.Now().UTC())
			return
		}
	}
}
