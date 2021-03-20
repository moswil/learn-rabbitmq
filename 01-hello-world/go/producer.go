package main

import (
	"github.com/streadway/amqp"
)

func main() {
	ch, q, err := ConnectToServer()

	body := "Hello World!"
	err = ch.Publish(
		"",     //exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	FailOnError(err, "Failed to publish a message")
}
