package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Service interface {
	Connect() error
}

type RabbitMQ struct {
	Conn *amqp.Connection
}

func (r *RabbitMQ) Connect() error {
	fmt.Println("Try connecting to Rabbitmq")
	//var err error
	_, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Successfully")
	return nil
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
