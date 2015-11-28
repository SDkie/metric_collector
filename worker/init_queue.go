package worker

import (
	"os"

	"github.com/SDkie/metric_collector/logger"
	"github.com/streadway/amqp"
)

const (
	Q_DISTINCT_NAME = "distinct_name"
	Q_HOURLY_LOG    = "hourly_log"
	Q_ACCOUNT_NAME  = "account_name"
)

var RabbitConnection *amqp.Connection
var RabbitChannel *amqp.Channel

func InitRabbitMQ() {
	var err error

	RabbitConnection, err = amqp.Dial(os.Getenv("RABBITMQ_URI"))
	logger.PanicfIfError(err, "Error while dialing to RabbitMQ Server, %s", err)

	RabbitChannel, err = RabbitConnection.Channel()
	logger.PanicfIfError(err, "Error while creating RabbitMQ Channel, %s", err)

	err = RabbitChannel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	logger.PanicfIfError(err, "Error while setting Qos paramter for RabbitMQ, %s", err)

	declareQueue(Q_ACCOUNT_NAME)
	declareQueue(Q_DISTINCT_NAME)
	declareQueue(Q_HOURLY_LOG)

	logger.Info("Rabbitmq Successfully Initialize")
}

func declareQueue(qName string) {
	_, err := RabbitChannel.QueueDeclare(
		qName, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	logger.PanicfIfError(err, "Error while declaring Queue - %s, %s", qName, err)
}
