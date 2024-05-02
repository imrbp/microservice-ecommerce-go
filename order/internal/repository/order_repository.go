package repository

import (
	"order_service/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderRepository struct {
  DB *gorm.DB
  Log *logrus.Logger
}

func NewOrderRepository(db *gorm.DB, log *logrus.Logger) *OrderRepository{
  return &OrderRepository{
    DB : db,
    Log : log,
  }
}

func (oR *OrderRepository) GetByFilter(searchFilter string, orders *[]entity.Orders) error {
  if err := oR.DB.Scopes(OrderStatus(searchFilter)).Find(&orders).Error ; err != nil {
    oR.Log.Panicf("Error Get By Filter (Order) : %v", searchFilter)
    return err
  }
  return nil
}

func OrderStatus(status string) func (tx *gorm.DB) *gorm.DB {
  return func (tx *gorm.DB) *gorm.DB  {
    if status == "complete" {
     tx = tx.Where("status = complete") 
    }
    if status == "shipping" {
     tx = tx.Where("status = shipping") 
    }
    if status == "pending" {
     tx = tx.Where("status = pending") 
    }
    if status == "cancel" {
     tx = tx.Where("status = cancel") 
    }
    return tx
  }
}

func (oR *OrderRepository) GetById(orderId string, order *entity.Orders) error {

  oR.Log.Infof("Get All (Order): %v", orderId)
  return oR.DB.Where("order_id = ?", order.OrderId).First(order).Error
}

func (oR *OrderRepository) CreateOrder(orders entity.Orders) error {

  oR.Log.Infof("Create (Order): %v", orders)
  return oR.DB.Create(orders).Error
}

func (oR *OrderRepository) UpdateOrder(payload *entity.Orders) error {

  oR.Log.Infof("Update (Order): %v", payload)
  return oR.DB.Where("order_id= ?", payload.OrderId).Save(payload).Error
}

