package converter

import (
	"order_service/internal/entity"
	"order_service/internal/model"
)

func ProductToResponse(product *entity.Product) *model.ProductResponse {
  return &model.ProductResponse{
    ProductId : product.ProductId,
    Title : product.Title,
    Description : product.Description,
    Price : product.Price, 
  }
}
