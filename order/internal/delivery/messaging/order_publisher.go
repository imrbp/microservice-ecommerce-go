package delivery

import (
	"encoding/json"
	"order_service/internal/entity"
	"order_service/internal/model"

	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type OrderPublisher struct {
	Logger         *logrus.Logger
	Connection     *rabbitmq.Conn
	OrderPublisher *rabbitmq.Publisher
}

func NewOrderPublisher(logger *logrus.Logger, conn *rabbitmq.Conn) *OrderPublisher {
	orderPublisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsExchangeName("proceed_payment"),
	)

	if err != nil {
		logger.Panicf("Error Make Order Cosumer: %v", err)
	}

	return &OrderPublisher{
		Logger:         logger,
		Connection:     conn,
		OrderPublisher: orderPublisher,
	}
}

func (publisher *OrderPublisher) Checkout(order *entity.Orders) error {
	model := model.OrderPayment{
		OrderId:    order.OrderId,
		TotalPrice: order.TotalPrice,
	}
	payload, err := json.Marshal(model)
	if err != nil {
		return err
	}
	err = publisher.OrderPublisher.Publish(
		payload,
		[]string{""},
		rabbitmq.WithPublishOptionsContentType("json"),
	)

	if err != nil {
		return err
	}

	return nil
}
