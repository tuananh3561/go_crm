package job

import (
	"github.com/streadway/amqp"
	"log"
	"reflect"
)

// Queue type
type Queue struct {
	channelRabbitMQ *amqp.Channel
}

// QueuePublish type
type QueuePublish struct {
	channelRabbitMQ *amqp.Channel
}

var queue *Queue

func DeclareToQueue(channelRabbitMQ *amqp.Channel, queueName string) {
	q := Queue{
		channelRabbitMQ: channelRabbitMQ,
	}
	queue = &q

	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err := channelRabbitMQ.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		panic(err)
	}
}

func SubscribingToQueue(channelRabbitMQ *amqp.Channel, queueName string) {
	q := Queue{
		channelRabbitMQ: channelRabbitMQ,
	}
	queue = &q

	// Subscribing to Queue for getting messages.
	messages, err := channelRabbitMQ.Consume(
		queueName, // queue name
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		log.Println(err)
	}

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", message.Body)
			Execute(q, "InsertLogMongo")
		}
	}()

	<-forever
}

func Execute(q Queue, seedMethodName string) {
	// Get the reflection value of the method
	m := reflect.ValueOf(q).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// Execute the method
	m.Call(nil)
}

func Publish(message amqp.Publishing) {
	errQ := queue.channelRabbitMQ.Publish(
		"",       // exchange
		"go_crm", // queue name
		false,    // mandatory
		false,    // immediate
		message,  // message to publish
	)
	if errQ != nil {
		log.Fatal("error publishing queue ", errQ)
	}
}
