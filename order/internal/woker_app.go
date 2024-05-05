package internal

import (
	delivery "order_service/internal/delivery/messaging"
	dependency "order_service/internal/depedency"
	"order_service/internal/gateway/messaging"
	"order_service/internal/repository"
	"order_service/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/wagslane/go-rabbitmq"
	"gorm.io/gorm"
)

type BootstrapWorkerApp struct {
  DB *gorm.DB
  App *fiber.App
  Log *dependency.Logger
  Validation *dependency.CustomValidator
  Conn *rabbitmq.Conn
}

func BootstrapWorker(config *BootstrapWorkerApp) {
  orderRepository := repository.NewOrderRepository(config.DB, config.Log.Log)

  orderPublisher := delivery.NewOrderPublisher(config.Log.Log, config.Conn)  
  orderConsumer := messaging.NewOrderConsumer(config.Log.Log, config.Conn) 


  orderPubSub := service.NewOrderPubSub(config.Log.Log, config.DB, orderRepository, orderPublisher, orderConsumer)

  if err := orderPubSub.Start() ; err != nil {
    config.Log.Log.Warnf("failed to start Consumer: %v", err)
  }
  
}
