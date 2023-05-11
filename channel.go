package amqp

import (
	"github.com/rabbitmq/amqp091-go"
)

type channel struct {
	*Connection
	*amqp091.Channel
}

type BrokerOption struct {
	Queue    *QueueOption
	Exchange *ExchangeOption
	Bind     *QueueBindOption
	Worker   *WorkerOption
}

type BrokerOpt func(*BrokerOption)

func WithQueueOption(opt *QueueOption) BrokerOpt {
	return func(broker *BrokerOption) {
		broker.Queue = opt
	}
}
func WithExchangeOption(opt *ExchangeOption) BrokerOpt {
	return func(broker *BrokerOption) {
		broker.Exchange = opt
	}
}
func WithBindOption(opt *QueueBindOption) BrokerOpt {
	return func(broker *BrokerOption) {
		broker.Bind = opt
	}
}
func WithWorkerOption(opt *WorkerOption) BrokerOpt {
	return func(broker *BrokerOption) {
		broker.Worker = opt
	}
}

type QueueOption struct {
	Name       string
	Durable    bool
	Exclusive  bool
	AutoDelete bool
	NoWait     bool
	Arguments  amqp091.Table
	MsgHeaders amqp091.Table
}

type ExchangeOption struct {
	Name       string
	Type       ExchangeType
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Arguments  amqp091.Table
}

type QueueBindOption struct {
	RouteKey  string
	NoWait    bool
	Arguments amqp091.Table
}

// WorkerOption 重试次数
type WorkerOption struct {
	ProuderRetryNumb int8
	ConsumerRetryNUm int8
}

func MakeDefaultBroker(queue string, exchange string, exchangeType ExchangeType, routeKey string) *BrokerOption {
	broker := &BrokerOption{}
	broker = broker.WithBrokerOpt(WithQueueOption(&QueueOption{
		Name:       queue,
		Durable:    true,
		Exclusive:  false,
		AutoDelete: false,
		NoWait:     false,
		Arguments:  nil,
	}), WithExchangeOption(&ExchangeOption{
		Name:       exchange,
		Type:       exchangeType,
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Arguments:  nil,
	}), WithBindOption(&QueueBindOption{
		RouteKey:  routeKey,
		NoWait:    false,
		Arguments: nil,
	}), WithWorkerOption(&WorkerOption{
		ProuderRetryNumb: 3,
		ConsumerRetryNUm: 3,
	}))

	return broker
}

func (broker *BrokerOption) WithBrokerOpt(opt ...BrokerOpt) *BrokerOption {
	for _, o := range opt {
		o(broker)
	}
	return broker
}

func (c *channel) Declare(broker *BrokerOption) (err error) {
	err = c.Queue(broker)
	if err != nil {
		return
	}
	err = c.Exchange(broker)
	if err != nil {
		return
	}
	err = c.Bind(broker)
	return
}

func (c *channel) Queue(broker *BrokerOption) (err error) {
	_, err = c.QueueDeclare(broker.Queue.Name, broker.Queue.Durable, broker.Queue.AutoDelete, broker.Queue.Exclusive, broker.Queue.NoWait, broker.Queue.Arguments)
	return
}

func (c *channel) Exchange(broker *BrokerOption) (err error) {
	err = c.ExchangeDeclare(broker.Exchange.Name, string(broker.Exchange.Type), broker.Exchange.Durable, broker.Exchange.AutoDelete, broker.Exchange.Internal, broker.Exchange.NoWait, broker.Exchange.Arguments)
	return
}

func (c *channel) Bind(broker *BrokerOption) (err error) {
	err = c.QueueBind(broker.Queue.Name, broker.Bind.RouteKey, broker.Exchange.Name, broker.Bind.NoWait, broker.Bind.Arguments)
	return
}

// wait前必须设置Confirm
func (c *channel) Confirm(broker *BrokerOption) error {
	return c.Channel.Confirm(broker.Exchange.NoWait)
}
