package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Service
type Service interface {
	Connect() error
	Publish(ctx context.Context, message string) error
	Consume()
}

// RabbitMQ
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func (r *RabbitMQ) Connect() error {
	fmt.Println("Try connecting to Rabbitmq")
	var err error
	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Successfully")

	r.Channel, err = r.Conn.Channel()
	if err != nil {
		return err
	}

	_, err = r.Channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	return nil
}

// Publish - take in a string message and publish to queue
func (r *RabbitMQ) Publish(ctx context.Context, message string) error {

	err := r.Channel.PublishWithContext(
		ctx,
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	fmt.Println("Successfully published to queue")

	return nil

}

func (r *RabbitMQ) Consume() {
	msgs, err := r.Channel.Consume(
		"TestQueue",
		"",
		true, false, false, false, nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	for msg := range msgs {
		fmt.Printf("Received Message : %s\n", msg.Body)
	}

}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
