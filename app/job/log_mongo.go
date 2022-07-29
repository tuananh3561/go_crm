package job

import (
	"github.com/streadway/amqp"
	"log"
)

func PublishInsertLogMongo() {
	log.Println("Publish InsertLogMongo")

	// Create a message to publish.
	messageQ := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(""),
	}

	// Attempt to publish a message to the queue.
	Publish(messageQ)
}

func (s Queue) InsertLogMongo() {
	log.Println("InsertLogMongo")
}
