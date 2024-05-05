package model

type CartResponse struct {
	Product  ProductResponse `json:"product"`
	Quantity int             `json:"quantity"`
}

type AddToCartRequest struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CartRequest struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
