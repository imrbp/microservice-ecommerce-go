package entity

type Orders struct {
  OrderId string `gorm:"primaryKey;not null"`
  Status string `gorm:"column:status;not null"`
  Airwaybill string `gorm:"column:airwaybill"`
  Invoice string `gorm:"column:invoice"`
  TotalPrice int64 `gorm:"column:total_price;not null"`
}
