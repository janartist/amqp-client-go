package amqp

import (
	"github.com/rabbitmq/amqp091-go"
	"time"
)

// SetDelayQueue 设置延迟队列
func (broker *BrokerOption) SetDelayQueue(t time.Duration) {
	if broker.Queue.MsgHeaders == nil {
		broker.Queue.MsgHeaders = amqp091.Table{}
	}
	//毫秒
	broker.Queue.Arguments["x-delay"] = t.Milliseconds()

	if broker.Exchange.Arguments == nil {
		broker.Exchange.Arguments = amqp091.Table{}
	}
	broker.Exchange.Arguments["x-delayed-type"] = "topic"
	broker.Exchange.Type = Delay
}
