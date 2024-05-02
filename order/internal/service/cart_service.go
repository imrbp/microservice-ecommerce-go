package service

import (
	"context"
	"order_service/internal/entity"
	"order_service/internal/model"
	"order_service/internal/model/converter"
	"order_service/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CartService struct {
  DB *gorm.DB
  Log *logrus.Logger
  CartRepository *repository.CartRepository
} 

func NewCartService(
  db *gorm.DB,
  log *logrus.Logger,
  cartRepository *repository.CartRepository,
) *CartService {
  return &CartService{
    DB : db,
    Log : log,
    CartRepository : cartRepository,
  }
}

func (cS *CartService) AddQuantityProduct(ctx context.Context, cartProductId string, quantity int) error {
  tx := cS.DB.WithContext(ctx).Begin()
  defer tx.Rollback()
 
  var cart entity.Cart

  if err := cS.CartRepository.GetById(cartProductId, &cart) ; err != nil {
    return err
  }


  cart.Quantity = quantity

  if err := cS.CartRepository.UpdateQuantityProduct(&cart) ; err != nil {
    return err
  }
  return nil
}


func (cS *CartService) RemoveProduct(ctx context.Context, productId string) error {
  tx := cS.DB.WithContext(ctx).Begin()
  defer tx.Rollback()

  var cart entity.Cart
  if err := cS.CartRepository.GetById(productId, &cart) ; err != nil {
    return err
  }

  if err := cS.CartRepository.RemoveProduct(&cart) ;err != nil {
    return err
  }

  return nil 
}

func (cS *CartService) GetAll(ctx context.Context, cartResponses *[]model.CartResponse) error {
  var cart []entity.Cart

  if err := cS.CartRepository.GetAll(&cart) ; err != nil {
    return err
  }

  for _, item := range cart {
    *cartResponses = append(*cartResponses, *converter.CartToResponse(&item))
  }
  return nil
} 
