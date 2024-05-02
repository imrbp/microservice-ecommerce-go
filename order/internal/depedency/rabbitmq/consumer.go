package rabbitmq

import (
	"order_service/internal/model"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
  Channel *ChannelManager
  ReconnectErrCh <-chan error
  CloseConnectionToManagerCh chan<- struct{}
  ConsumerModel model.ConsumerModel
}

type Action int

type Handler func(d Delivery) (action Action)

type Delivery struct {
  amqp.Delivery
}

