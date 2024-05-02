package repository

import (
	"order_service/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CartRepository struct {
  DB *gorm.DB
  Log *logrus.Logger
}

func NewCartRepository(db *gorm.DB, log *logrus.Logger) *CartRepository{
  return &CartRepository{
    DB : db,
    Log : log,
  }
}

func (cR *CartRepository) AddToCart(productCart *entity.Cart) error {
  if err := cR.DB.FirstOrCreate(&productCart, entity.Cart{ProductId: productCart.ProductId}).Error ; err != nil {
    cR.Log.Warnf("Error Add (Cart): %v", err) 
    return err
  }
  cR.Log.Infof("Insert (Cart): %v", productCart)
  return nil
}

func (cR *CartRepository) UpdateQuantityProduct(payload *entity.Cart) error {
  cR.Log.Infof("Update (Cart): %v", payload) 
  return cR.DB.Where("product_id = ?", payload.ProductId).Save(payload).Error
}

func (cR *CartRepository) RemoveProduct(cart *entity.Cart) error {
  cR.Log.Infof("Remove (Cart) : %v", cart)
  return cR.DB.Where("product_id = ?", cart.ProductId).Delete(&cart).Error
}

func (cR *CartRepository) GetById(productId string, cart *entity.Cart) error {
  cR.Log.Infof("Get By Id (Cart): %v", productId)
  return cR.DB.Where("product_id = ? ", productId).First(&cart).Error 
} 

func(cR *CartRepository) GetAll(carts *[]entity.Cart) error{
  if err := cR.DB.Preload(clause.Associations).Find(&carts).Error; err != nil {
    cR.Log.Warnf("Error Get All: %v", err) 
    return err
  }
  cR.Log.Infof("Get All (Cart): %v", carts) 
  return nil
}
