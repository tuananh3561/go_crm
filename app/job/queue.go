package job

import (
	"github.com/tuananh3561/go_crm/app/config"
	"github.com/tuananh3561/go_crm/app/job/rabbit_mq"
	"log"
	"reflect"
)

type Queue struct {
	cf config.Config
}

func RunQueue(cf config.Config, queueNames ...string) {
	queue := Queue{
		cf: cf,
	}

	queueName := cf.Queue.Queue.Name
	if len(queueNames) > 0 {
		queueName = queueNames[0]
	}

	// Subscribing to Queue for getting messages.
	messages := rabbit_mq.SubscribingToQueue(cf.Queue.ChannelRabbitMQ, queueName)

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", message)
			log.Printf(" > Received message: %s\n", message.Body)
			Execute(queue, "InsertLogMongo", message.Body)
		}
	}()

	<-forever
}

func Execute(q Queue, seedMethodName string, body []byte) {
	// Get the reflection value of the method
	m := reflect.ValueOf(q).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// Execute the method
	m.Call(nil)
}
