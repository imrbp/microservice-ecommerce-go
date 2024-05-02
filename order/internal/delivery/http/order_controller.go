package http

import (
	"context"
	dependency "order_service/internal/depedency"
	"order_service/internal/model"
	"order_service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
  Validation *dependency.CustomValidator
  OrderService *service.OrderService
}

func NewOrderController(validation *dependency.CustomValidator, orderService *service.OrderService) *OrderController {
  return &OrderController{
    Validation : validation,
    OrderService : orderService,
  }
}

func (c *OrderController) GetAll(ctx *fiber.Ctx) error {
  request := new(model.SearchOder)

  if err := c.Validation.ParseBody(ctx, &request) ; err != nil {
    return err
  }

  if err := c.OrderService.GetAllOrder(context.Background(), request) ; err != nil {
    return err
  }

  c.OrderService.Log.Infof("Get All Orders: %v", ctx.BaseURL())
 

  return ctx.Status(fiber.StatusOK).JSON(
    request,
  )
}


func (c *OrderController) Checkout(ctx *fiber.Ctx) error {
  var checkoutRequest model.OrderRequest
  
  if err := c.Validation.ParseBody(ctx, &checkoutRequest) ; err != nil {
    return err
  }
  
  if err := c.OrderService.Checkout(context.Background(), &checkoutRequest) ; err != nil {
    return err
  }
  c.OrderService.Log.Infof("Checkout Orders: %v", ctx.BaseURL())
 
  return nil

}
