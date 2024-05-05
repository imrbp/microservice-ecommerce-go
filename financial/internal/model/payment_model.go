package model

type OrderPayment struct {
	OrderId       string `json:"order_id"`
	TotalPrice    int64  `json:"total_price"`
}
