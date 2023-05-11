package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	rabbitmq "rabbitmq-client"
)

var (
	uri = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
)

func init() {
	flag.Parse()
}

func main() {
	procuder()
	procuderWithPool()

}

func procuder() {
	conn, _ := rabbitmq.NewConnection(*uri)
	ch, err := conn.NewChannel()

	if err != nil {
		fmt.Printf("channel error:%v \n", err)
		return
	}
	broker := rabbitmq.MakeDefaultBroker("default-queue", "default-exchange", rabbitmq.Direct, "default")
	ch.Declare(broker)
	publish := make(chan string)

	for i := 0; i < 100; i++ {
		go func() {
			body := time.Now().String()
			_, err := ch.Publish(context.Background(), broker, body)
			if err != nil {
				fmt.Printf("publish error:%v \n", err)
				return
			}
			publish <- body
		}()
	}

	for {
		select {
		case no := <-publish:
			fmt.Printf("publish success: %s \n", no)
		}
	}
}

func procuderWithPool() {
	amqp, err := rabbitmq.NewRabbitmq(&rabbitmq.Rabbitmq{
		Uri:        *uri,
		InitialCap: 5,
		MaxCap:     20,
	})
	if err != nil {
		fmt.Printf("pool amqp error:%v \n", err)
		return
	}
	conn, _ := amqp.GetConnection()
	ch, err := conn.NewChannel()

	if err != nil {
		fmt.Printf("pool channel error:%v \n", err)
		return
	}

	broker := rabbitmq.MakeDefaultBroker("default-queue", "default-exchange", rabbitmq.Direct, "default")

	err = ch.Confirm(broker)
	if err != nil {
		fmt.Printf("pool channel error:%v \n", err)
		return
	}
	ch.Declare(broker)
	publish := make(chan string)

	for i := 0; i < 100; i++ {
		go func() {
			body := time.Now().String()
			_, err := ch.Publish(context.Background(), broker, body)
			if err != nil {
				fmt.Printf("pool publish error:%v \n", err)
				return
			}
			publish <- body
		}()
	}

	for {
		select {
		case no := <-publish:
			fmt.Printf("pool publish success: %s \n", no)
		}
	}
}
