package converter

import (
	"order_service/internal/entity"
	"order_service/internal/model"
)

func OrderToResponse(order *entity.Orders) *model.OrderResponse {
  return &model.OrderResponse{
    OrderId : order.OrderId,
    Status : order.Status,
    Airwaybill : order.Airwaybill,
    InvoiceNumber : order.Invoice,
    TotalPrice : order.TotalPrice,
  }
}

func M_OrderProduct_E_OrderProduct(orderId string, product *model.OrderProductRequest) *entity.OrderProducts {
  return &entity.OrderProducts{
    OrderId : orderId,
    ProductId : product.ProductId,
    Quantity : product.Quantity,
    Price : product.Price,
  }
}
