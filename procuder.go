package amqp

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
)

func (c *channel) Publish(ctx context.Context, broker *BrokerOption, body string) (*amqp091.DeferredConfirmation, error) {
	confirm, err := c.PublishWithDeferredConfirmWithContext(
		ctx,
		broker.Exchange.Name, // publish to an exchange
		broker.Bind.RouteKey, // routing to 0 or more queues
		false,                // mandatory
		false,                // immediate
		amqp091.Publishing{
			Headers:         broker.Queue.MsgHeaders,
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp091.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,                 // 0-9
			// a bunch of application/implementation-specific fields
		},
	)
	return confirm, err
}
