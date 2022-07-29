package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/tuananh3561/go_crm/app/config"
	"github.com/tuananh3561/go_crm/app/database/migrations"
	"github.com/tuananh3561/go_crm/app/database/seeds"
	"github.com/tuananh3561/go_crm/app/job"
	"github.com/tuananh3561/go_crm/routes"
	"os"
)

func main() {
	cf := config.SetupConfig()

	defer config.CloseDatabaseConnection(cf.Database)
	//defer config.CloseAMQP()

	handleArgs(cf)

	handleRun(cf)
}

func handleRun(cf config.Config) {
	service := gin.Default()
	config.SetUseSentry(service)

	routes.Router(service, cf.Database, cf.ChannelRabbitMQ)

	err := service.Run(":" + cf.App.AppPort)
	if err != nil {
		panic("Run service failed")
	}
}

func handleArgs(cf config.Config) {
	flag.Parse()
	args := flag.Args()
	if len(args) >= 1 {
		switch args[0] {
		case "migration":
			migrations.Migrations(cf.Database.MysqlAuth)
			os.Exit(0)
		case "seed":
			seeds.Execute(cf.Database, args[1:]...)
			os.Exit(0)
		case "queue":
			job.SubscribingToQueue(cf.ChannelRabbitMQ)
			os.Exit(0)
		}
	}
}
