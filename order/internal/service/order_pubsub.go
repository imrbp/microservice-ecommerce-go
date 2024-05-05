package service

import (
	"encoding/json"
	pub "order_service/internal/delivery/messaging"
	"order_service/internal/entity"
	con "order_service/internal/gateway/messaging"
	"order_service/internal/model"
	"order_service/internal/repository"

	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
	"gorm.io/gorm"
)

type OrderPubSub struct {
	Logger *logrus.Logger
	DB     *gorm.DB

	OrderRepository *repository.OrderRepository

	OrderPublisher *pub.OrderPublisher
	OrderConsumer  *con.OrderConsumer
}

func NewOrderPubSub(logger *logrus.Logger, db *gorm.DB, orderRepository *repository.OrderRepository, orderPublisher *pub.OrderPublisher, orderConsumer *con.OrderConsumer) *OrderPubSub {
	return &OrderPubSub{
		Logger:          logger,
		DB:              db,
		OrderRepository: orderRepository,
		OrderPublisher:  orderPublisher,
		OrderConsumer:   orderConsumer,
	}
}

func (pubSub *OrderPubSub) UpdateStatus(payload []byte) error {
	order := model.UpdateOrder{}

	err := json.Unmarshal(payload, &order)

	if err != nil {
		return err
	}

	orderUpdaate := entity.Orders{
		OrderId: order.OrderId,
		Status:  order.Status,
		Invoice: order.InvoiceNumber,
	}
	if err := pubSub.OrderRepository.UpdateOrder(&orderUpdaate); err != nil {
		return err
	}
	return nil
}

func (pubSub *OrderPubSub) Start() error {
	err := pubSub.OrderConsumer.Consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
		err := pubSub.UpdateStatus(d.Body)
		if err != nil {
			pubSub.Logger.Warnf("Error when update Order : %v", d.Body)
		}
		return rabbitmq.Ack
	})

	if err != nil {
		return err
	}

	return nil
}
