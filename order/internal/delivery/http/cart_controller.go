package http

import (
	"context"
	dependency "order_service/internal/depedency"
	"order_service/internal/model"
	"order_service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CartController struct {
  Validation *dependency.CustomValidator
  CartService *service.CartService
}

func NewCartController(validation *dependency.CustomValidator, cartService *service.CartService) *CartController {
  return &CartController{
    Validation : validation,
    CartService : cartService,
  }
}

func (c *CartController) AddQuantityProduct(ctx *fiber.Ctx) error {
  var cartRequest model.CartRequest

  if err := c.Validation.ParseBody(ctx, &cartRequest) ; err != nil {
    return err
  }

  if err := c.CartService.AddQuantityProduct(context.Background(), cartRequest.ProductId, cartRequest.Quantity); err != nil {
    return err
  }

  c.CartService.Log.Infof("POST add quantity : %v", ctx.BaseURL())
  return ctx.Status(fiber.StatusOK).JSON(
    "Success",
    )
}

func (c *CartController) DeleteProduct(ctx *fiber.Ctx) error {
  var cartRequest model.CartRequest

  if err := c.Validation.ParseBody(ctx, &cartRequest) ; err != nil {
    return err
  }

  if err := c.CartService.RemoveProduct(context.Background(), cartRequest.ProductId) ; err != nil {
    return err
  }

  c.CartService.Log.Infof("Delete Product : %v", ctx.BaseURL())
 
  return ctx.Status(fiber.StatusOK).JSON(
    "success",
    )
}

func (c *CartController) GetAll(ctx *fiber.Ctx) error {
  var cartsRequest []model.CartResponse

  if err := c.CartService.GetAll(context.Background(), &cartsRequest) ; err != nil {
    return err
  }

  c.CartService.Log.Infof("Get All Products : %v", ctx.BaseURL())
 
  return ctx.Status(fiber.StatusOK).JSON(
    cartsRequest,
    )
}
