package rabbit_mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type ConfigRabbitMQ struct {
	Url       string
	QueueName string
}

type QueueRabbitMQ struct {
	ConnectRabbitMQ *amqp.Connection
	ChannelRabbitMQ *amqp.Channel
	Queue           amqp.Queue
}

var ChannelRabbitMQ *amqp.Channel
var Queue amqp.Queue

func InitRabbitMQ(config ConfigRabbitMQ) QueueRabbitMQ {
	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(config.Url)
	if err != nil {
		panic(err)
	}

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	ChannelRabbitMQ = channelRabbitMQ
	//defer channelRabbitMQ.Close()

	queue := DeclareToQueue(channelRabbitMQ, config.QueueName)
	Queue = queue

	return QueueRabbitMQ{connectRabbitMQ, channelRabbitMQ, queue}
}

func CloseRabbitMQ(queue QueueRabbitMQ) {
	errConnect := queue.ConnectRabbitMQ.Close()
	if errConnect != nil {
		panic("Failed to close database connect RabbitMQ")
	}
	errChannel := queue.ChannelRabbitMQ.Close()
	if errChannel != nil {
		panic("Failed to close database channel RabbitMQ")
	}
}

func DeclareToQueue(channelRabbitMQ *amqp.Channel, queueName string) amqp.Queue {
	// With the instance and declare Queues that we can
	// publish and subscribe to.
	q, err := channelRabbitMQ.QueueDeclare(
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
	return q
}

func SubscribingToQueue(channelRabbitMQ *amqp.Channel, queueName string) <-chan amqp.Delivery {
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
	return messages
}

func Publish(data interface{}) {
	dataByte, _ := json.Marshal(data)
	// Create a message to publish.
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        dataByte,
	}
	// Attempt to publish a message to the queue.
	errQ := ChannelRabbitMQ.Publish(
		"",         // exchange
		Queue.Name, // queue name
		false,      // mandatory
		false,      // immediate
		message,    // message to publish
	)
	if errQ != nil {
		log.Fatal("error publishing queue ", errQ)
	}
}
