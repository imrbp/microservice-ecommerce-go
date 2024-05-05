package service

import (
	"financial_service/internal/repository"
  con "financial_service/internal/gateway/messaging"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PaymentPubCon struct {
  Logger *logrus.Logger
  DB *gorm.DB

  PaymentRepository *repository.PaymentRepository
  PaymnetConsumer *con.PaymentConsumer
}

func NewPaymentPubCon(logger *logrus.Logger, db *gorm.DB, paymentRepository *repository.PaymentRepository, paymentConsumer *con.PaymentConsumer) *PaymentPubCon {
  return &PaymentPubCon{
    Logger : logger,
    DB : db,
    PaymentRepository : paymentRepository,
  }
}

