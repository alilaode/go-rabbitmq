package main

import (
	"context"
	"fmt"
	"time"

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
	defer app.rmq.Conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = app.rmq.Publish(ctx, "Hi")
	if err != nil {
		return err
	}

	app.rmq.Consume()

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error setting up")
		fmt.Println(err)
	}
}
