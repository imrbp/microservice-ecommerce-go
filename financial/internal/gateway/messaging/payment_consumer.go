package con

import (
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type PaymentConsumer struct {
  Consumer *rabbitmq.Consumer
}

func NewPaymentConsumer(logger *logrus.Logger,  conn *rabbitmq.Conn) *PaymentConsumer {
  consumer, err := rabbitmq.NewConsumer(conn, "order_service")
  if err != nil {
    logger.Panicf("Error when deploy new Consumer : %v", err)
  }
  return &PaymentConsumer{
    Consumer : consumer,
  }
}
