package model

type OrderResponse struct {
  OrderId string `json:"order_id"`
  Status string `json:"status"`
  Airwaybill string `json:"airwaybill"`
  InvoiceNumber string `json:"invoice_number"`
  TotalPrice  int64 `json:"total_price"`
}

type OrdeDetailResponse struct {
  Order OrderResponse `json:"order"`
  Products []OrderProductResponse `json:"products"`
}

type OrderProductResponse struct {
  ProductId string `json:"product_id"`
  Title string `json:"title"`
  Quantity int `json:"quantity"`
  Price int64 `json:"price"`
}

type OrderProductRequest struct {
  ProductId string `json:"product_id" validate:"uuid,required"` 
  Quantity int `json:"quantity" validate:"gte=0,required"`
  Price int64 
}

type OrderRequest struct {
  Products []OrderProductRequest `json:"products"`
}

type UpdateOrder struct {
  OrderId string `json:"order_id"`
  Status string `json:"status"`
  Airwaybill string `json:"airwaybill"`
  InvoiceNumber string `json:"invoice_number"`
}

type SearchOder struct {
  FilterBy string `json:"filter_by"` 
  Orders []OrderResponse `json:"orders"`
}
