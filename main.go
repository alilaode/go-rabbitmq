package main

import (
	"fmt"

	"github.com/alilaode/go-rabbitmq/internal/rabbitmq"
)

type App struct {
	rmq *rabbitmq.RabbitMQ
}

func Run() error {
	app := App{
		rmq: rabbitmq.NewRabbitMQService(),
	}

	err := app.rmq.Connect()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error setting up")
		fmt.Println(err)
	}
}
