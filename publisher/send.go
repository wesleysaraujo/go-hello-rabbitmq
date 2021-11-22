package publisher

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Message struct {
	Message string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatal("%s: %s", msg, err)
	}
}

func Send() {
	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to ope a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name of queue
		false, // is durable
		false, // delete when unused
		false, // exlcusive
		false, // no-wait
		nil, // arguments
	)
	failOnError(err, "Failed to declare a queue")

	message := Message {
		Message: "Hello World",
	}

	body, err := json.Marshal(message)
	failOnError(err, "Failed format message in json")
	err = ch.Publish(
		"",  // exchange
		q.Name, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing {
			ContentType: "application/json",
			Body: []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}