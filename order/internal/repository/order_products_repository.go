package repository

import (
	"order_service/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderProductsRespository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

func NewOrderProductsRespository(db *gorm.DB, log *logrus.Logger) *OrderProductsRespository {
	return &OrderProductsRespository{
		DB:  db,
		Log: log,
	}
}

func (oR *OrderProductsRespository) GetById(orderId string, orderProducts *[]entity.OrderProducts) error {
	if err := oR.DB.Where("order_id = ?", orderId).Find(&orderProducts).Error; err != nil {
		return err
	}
	return nil
}

func (oR *OrderProductsRespository) GetProductById(productId string, orderProduct *entity.OrderProducts) error {
	return oR.DB.Where("order_id = ?", productId).First(&orderProduct).Error
}

func (oR *OrderProductsRespository) CreateOrderProducts(orderProducts *[]entity.OrderProducts) error {
	return oR.DB.Create(orderProducts).Error
}
