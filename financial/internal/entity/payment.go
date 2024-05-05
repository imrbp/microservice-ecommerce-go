package entity

type Payment struct {
  Id string `gorm:"column:id"`
  InvoiceNumber string `gorm:"column:invoice_number"`
  OrderId string `gorm:"column:order_id"`
  Status string `gorm:"column:status"`
  Amount int64 `gorm:"column:amount"`
}
