package entity

type Product struct {
	ProductId   string `gorm:"column:product_id;primaryKey; not null"`
	Title       string `gorm:"column:title;not null"`
	Description string `gorm:"column:description;not null"`
	Price       int64  `gorm:"column:price;not null"`
}
