package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type DlxOption struct {
	DlxExchange   string
	DlxRouteKey   string
	DlxMessageTtl time.Duration
}

// SetDlxQueue 设置死信队列
func (broker *BrokerOption) SetDlxQueue(option *DlxOption) {
	if broker.Queue.Arguments == nil {
		broker.Queue.Arguments = amqp.Table{}
	}
	if option.DlxExchange == "" {
		option.DlxExchange = broker.Exchange.Name + ".dlx"
	}
	if option.DlxRouteKey == "" {
		option.DlxRouteKey = broker.Bind.RouteKey + ".dlx"
	}
	//毫秒
	broker.Queue.Arguments["x-message-ttl"] = option.DlxMessageTtl.Milliseconds()
	broker.Queue.Arguments["x-dead-letter-exchange"] = option.DlxExchange
	broker.Queue.Arguments["x-dead-letter-routing-key"] = option.DlxRouteKey
}
