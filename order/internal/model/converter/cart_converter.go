package converter

import (
	"order_service/internal/entity"
	"order_service/internal/model"
)

func CartToResponse(cart *entity.Cart) *model.CartResponse {
  return &model.CartResponse{
    Product : model.ProductResponse(cart.Product),
    Quantity : cart.Quantity,
  }
}

