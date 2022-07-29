package config

import (
	"github.com/streadway/amqp"
	"github.com/tuananh3561/go_crm/app/job"
)

type AMQPConfig struct {
	amqpServerURL string
}

var ConnectRabbitMQ *amqp.Connection
var ChannelRabbitMQ *amqp.Channel

func getConfigAMQP() AMQPConfig {
	var aMQPConfig = AMQPConfig{}
	// Define RabbitMQ server URL.
	aMQPConfig.amqpServerURL = GetEnv("AMQP_SERVER_URL", "")

	return aMQPConfig
}

func AMQPInit() *amqp.Channel {
	// Define RabbitMQ server config.
	var aMQPConfig = getConfigAMQP()

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(aMQPConfig.amqpServerURL)
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

	job.DeclareToQueue(channelRabbitMQ)

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
