package messaging

import (
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type OrderConsumer struct {
  Logger *logrus.Logger
  Connection *rabbitmq.Conn
  Consumer *rabbitmq.Consumer
}

func NewOrderConsumer(logger *logrus.Logger, conn *rabbitmq.Conn) *OrderConsumer {
	orderConsumer, err := rabbitmq.NewConsumer(
		conn,
		"order_service",
		rabbitmq.WithConsumerOptionsRoutingKey(""),
		rabbitmq.WithConsumerOptionsExchangeName("payment_status"),
		rabbitmq.WithConsumerOptionsExchangeKind("fanout"),
		rabbitmq.WithConsumerOptionsQueueDurable,
		rabbitmq.WithConsumerOptionsQueueQuorum,
	)
	if err != nil {
		logger.Panicf("Error Make Order Cosumer: %v", err)
	}

  return &OrderConsumer{
    Logger : logger,
    Connection : conn,
    Consumer : orderConsumer,
  }
}

