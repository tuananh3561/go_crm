package config

import (
	"github.com/tuananh3561/go_crm/app/job/rabbit_mq"
)

type Config struct {
	App      AppConfig
	Sentry   SentryConfig
	Queue    rabbit_mq.QueueRabbitMQ
	Database Database
}

func SetupConfig() Config {
	LoadEnv()
	app := SetupConfigApp()
	LoadEnvironment()
	SetupLogOutput()
	sentry := SentryInit()
	queue := SetupQueueInit()
	database := SetupDatabaseConnection()
	config := Config{
		App:      app,
		Sentry:   sentry,
		Queue:    queue,
		Database: database,
	}
	return config
}
