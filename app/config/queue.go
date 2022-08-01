package config

import (
	"github.com/tuananh3561/go_crm/app/job/rabbit_mq"
)

func getConfigRabbitMQ() rabbit_mq.ConfigRabbitMQ {
	configRabbitMQ := rabbit_mq.ConfigRabbitMQ{}
	// Define RabbitMQ server URL.
	configRabbitMQ.Url = GetEnv("AMQP_URL", "")
	configRabbitMQ.QueueName = GetEnv("QUEUE_NAME", "go_crm")

	return configRabbitMQ
}

func SetupQueueInit() rabbit_mq.QueueRabbitMQ {
	// Define RabbitMQ server config.
	config := getConfigRabbitMQ()

	queue := rabbit_mq.InitRabbitMQ(config)
	return queue
}

func CloseQueue(queue rabbit_mq.QueueRabbitMQ) {
	rabbit_mq.CloseRabbitMQ(queue)
}
