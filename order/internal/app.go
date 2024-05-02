package internal

import (
	delivery "order_service/internal/delivery/http"
	"order_service/internal/delivery/http/route"
	dependency "order_service/internal/depedency"
	"order_service/internal/repository"
	"order_service/internal/service"
  "github.com/wagslane/go-rabbitmq"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BootstrapApp struct {
  DB *gorm.DB
  App *fiber.App
  Log *dependency.Logger
  Validation *dependency.CustomValidator
  Conn *rabbitmq.Conn
}

func Bootstrap(config *BootstrapApp) {
  productRepository := repository.NewProductsRepository(config.DB, config.Log.Log)
  cartRepository := repository.NewCartRepository(config.DB, config.Log.Log)
  orderRepository := repository.NewOrderRepository(config.DB, config.Log.Log)
  orderProductRepository := repository.NewOrderProductsRespository(config.DB, config.Log.Log)

  orderPublisher, err := rabbitmq.NewPublisher(
    config.Conn,
    rabbitmq.WithPublisherOptionsExchangeName("proceed_payment"),
    )
  
  if err != nil {
    config.Log.Log.Panicf("Error Make Order Publisher : %v", err)
  }

  orderConsumer, err := rabbitmq.NewConsumer(
  	config.Conn,
  	"order_service",
    rabbitmq.WithConsumerOptionsRoutingKey(""),
    rabbitmq.WithConsumerOptionsExchangeName("payment_status"),
    rabbitmq.WithConsumerOptionsExchangeKind("fanout"),
    rabbitmq.WithConsumerOptionsQueueDurable,
    rabbitmq.WithConsumerOptionsQueueQuorum,
    )
  if err != nil {
    config.Log.Log.Panicf("Error Make Order Cosumer: %v", err)
  }


  productService := service.NewProductService(config.DB, config.Log.Log, productRepository, cartRepository)
  cartService := service.NewCartService(config.DB, config.Log.Log, cartRepository)
  orderService := service.NewOrderService(config.DB, config.Log.Log, orderRepository, cartRepository, productRepository, orderProductRepository, orderPublisher, orderConsumer )

  productController := delivery.NewProductController(config.Validation, productService)
  cartController := delivery.NewCartController(config.Validation, cartService)
  orderController := delivery.NewOrderController(config.Validation, orderService)

  routeApp := route.RouteConfig{
    App : config.App,
    ProductController : productController,
    CartController : cartController,
    OrderController : orderController,
  }
  routeApp.Setup()
}
