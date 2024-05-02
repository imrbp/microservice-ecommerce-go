package http

import (
	"context"
	dependency "order_service/internal/depedency"
	"order_service/internal/model"
	"order_service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
  Validation *dependency.CustomValidator
  ProductService *service.ProductService
}

func NewProductController(validation *dependency.CustomValidator, productService *service.ProductService) *ProductController {
  return &ProductController{
    Validation : validation,
    ProductService : productService,
  }
}

func (c *ProductController) GetAll(ctx *fiber.Ctx) error {
  request := new(model.SearchProduct)

  if err := c.Validation.ParseBody(ctx, &request) ; err != nil {
    return err
  }

  switch true {
  case request.Size <= 0:
    request.Size = 1;
    break;
  case request.CurrentPage <= 0:
    request.CurrentPage = 1
    break;
  }

  if err := c.ProductService.Get(context.Background(), ctx.Query("search"), request) ; err != nil {
    return err
  }

  return ctx.Status(fiber.StatusOK).JSON(
    request,
  )
}


func (c *ProductController) GetById(ctx *fiber.Ctx) error {
  productResponse :=  new(model.ProductResponse);

  param := struct {ProductId string `params:"productId"`}{}

  ctx.ParamsParser(&param)

  if err := c.ProductService.GetById(context.Background() ,param.ProductId, productResponse) ; err != nil {
    return err
  }

  return ctx.Status(fiber.StatusOK).JSON(
    productResponse,
    )
}

func (c *ProductController) AddToCart(ctx *fiber.Ctx) error {
  var cartResponse model.AddToCartRequest

  if err := c.Validation.ParseBody(ctx, &cartResponse) ; err != nil {
    return err
  }

  if err := c.ProductService.AddToCart(context.Background(), &cartResponse) ; err != nil {
    return err
  }

  return ctx.Status(fiber.StatusOK).JSON(
    "Success add",
    )
}
