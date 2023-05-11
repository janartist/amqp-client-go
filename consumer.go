package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumerOption struct {
	Queue       string
	ConsumerTag string
	NoLocal     bool
	AutoAck     bool
	Exclusive   bool
	NoWait      bool
	Arguments   amqp.Table
}
type DeliveryFunc func(amqp.Delivery) error

func (c *channel) BasicConsumer(option *ConsumerOption, fn DeliveryFunc) {
	deliveries, err := c.Consume(
		option.Queue,
		option.ConsumerTag,
		option.AutoAck,
		option.Exclusive,
		option.NoLocal,
		option.NoWait,
		option.Arguments,
	)
	if err != nil {
		fmt.Printf("cannot consume from: %s, %v", option.Queue, err)
		return
	}
	fmt.Printf("subscribed queue...[%s]", option.Queue)
	for d := range deliveries {
		err := fn(d)
		if err == nil && option.AutoAck {
			d.Ack(false)
		}
	}
}
