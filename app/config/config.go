package config

import (
	"github.com/streadway/amqp"
)

type Config struct {
	App             AppConfig
	Sentry          SentryConfig
	ChannelRabbitMQ *amqp.Channel
	Database        Database
}

func SetupConfig() Config {
	LoadEnv()
	app := SetupConfigApp()
	LoadEnvironment()
	SetupLogOutput()
	sentry := SentryInit()
	//var channelRabbitMQ = AMQPInit()
	database := SetupDatabaseConnection()
	config := Config{
		App:    app,
		Sentry: sentry,
		//ChannelRabbitMQ: channelRabbitMQ,
		Database: database,
	}
	return config
}
