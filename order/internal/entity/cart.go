package entity

type Cart struct {
  ProductId string 
	Product Product `gorm:"foreignKey:ProductId;references:ProductId;not null"`
	Quantity  int `gorm:"quantity;default:0;not null"`
}
