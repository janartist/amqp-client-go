package amqp

import "github.com/rabbitmq/amqp091-go"

// SetXHaPolicyALLArg 分布式集群高可用默认配置
func (broker *BrokerOption) SetXHaPolicyALLArg() {
	if broker.Queue.Arguments == nil {
		broker.Queue.Arguments = amqp091.Table{}
	}
	broker.Queue.Arguments["x-ha-policy"] = "all"
}
