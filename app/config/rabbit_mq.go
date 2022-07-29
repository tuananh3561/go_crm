package config

import (
	"github.com/streadway/amqp"
	"github.com/tuananh3561/go_crm/app/job"
)

type ConfigAMQP struct {
	AMQPUrl   string
	QueueName string
}

var ConfigAMQPDefault ConfigAMQP
var ConnectRabbitMQ *amqp.Connection
var ChannelRabbitMQ *amqp.Channel

func getConfigAMQP() ConfigAMQP {
	configAMQP := ConfigAMQP{}
	// Define RabbitMQ server URL.
	configAMQP.AMQPUrl = GetEnv("AMQP_URL", "")
	configAMQP.QueueName = GetEnv("QUEUE_NAME", "go_crm")

	return configAMQP
}

func AMQPInit() *amqp.Channel {
	// Define RabbitMQ server config.
	config := getConfigAMQP()
	ConfigAMQPDefault = config

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(config.AMQPUrl)
	if err != nil {
		panic(err)
	}
	ConnectRabbitMQ = connectRabbitMQ

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	ChannelRabbitMQ = channelRabbitMQ

	job.DeclareToQueue(channelRabbitMQ, config.QueueName)

	return channelRabbitMQ
}

func CloseAMQP() {
	errConnect := ConnectRabbitMQ.Close()
	if errConnect != nil {
		panic("Failed to close database connect RabbitMQ")
	}
	errChannel := ChannelRabbitMQ.Close()
	if errChannel != nil {
		panic("Failed to close database channel RabbitMQ")
	}
}
