package utils

import "github.com/streadway/amqp"

// ConnectToServer connect to the RabbitMQ server
func ConnectToServer() (*amqp.Channel, amqp.Queue, error) {
	// connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// create channel (where most of the API for getting things done resides)
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// to send a message we must declare a queue; then we can publish a message to the queue
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when used
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	return ch, q, err
}
