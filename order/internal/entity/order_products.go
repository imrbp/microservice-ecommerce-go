package entity

type OrderProducts struct {
  OrderId string `gorm:"foreignKey:order_id;references:order_id;not null"`
  ProductId string `gorm:"foreignKey:product_id;references:product_id;not null"`
  Quantity int `gorm:"column:quantity"`
  Price int64 `gorm:"column:price"`
}	
