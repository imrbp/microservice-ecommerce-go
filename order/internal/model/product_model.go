package model

type ProductResponse struct {
  ProductId string `json:"product_id"`
  Title string `json:"title"`
  Description string `json:"description"`
  Price int64 `json:"price"`
}

type ProductRequest struct {
  ProductId string `json:"product_id"`
  Title string
  Price int64
}

type SearchProduct struct {
  Size int `json:"size"`
  Total int64 `json:"total"`
  CurrentPage int `json:"current_page"`
  Products []ProductResponse `json:"products"`
} 
