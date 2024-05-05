package repository

import (
	"financial_service/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PaymentRepository struct {
  Logger *logrus.Logger
  DB *gorm.DB
}

func NewPaymentRepository(logger *logrus.Logger, db *gorm.DB) *PaymentRepository {
  return &PaymentRepository{
    Logger : logger,
    DB : db,
  }
}

func (r *PaymentRepository) Create(payment *entity.Payment) error {
  if err := r.DB.Where("payment_id = ?", payment.Id).FirstOrCreate(&payment).Error ; err != nil {
    return err
  }
  return nil
}

func (r *PaymentRepository) Update(payment *entity.Payment) error {
  if err := r.DB.Where("payment_id = ?", payment.Id).Save(&payment).Error ; err != nil{
    return err
  }
  return nil
} 
