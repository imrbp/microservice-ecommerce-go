package repository

import (
	"order_service/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepository struct {
  DB *gorm.DB
  Log *logrus.Logger
}

func NewProductsRepository(db *gorm.DB, log *logrus.Logger ) *ProductRepository {
  return &ProductRepository{
    DB : db,
    Log : log,
  }
}

func (pR *ProductRepository) GetAllProducts(searchParams string,size int, currentPage int, products *[]entity.Product) (int64, error) {
  var total int64
  if err := pR.DB.Scopes(pR.FilterProducts(searchParams)).Limit(size).Offset((currentPage - 1) * size).Find(&products).Error; err != nil {
    pR.Log.Warnf("Error Get All (Products) : %v", err)
    return total, err 
  }

  if err := pR.DB.Model(&entity.Product{}).Scopes(pR.FilterProducts(searchParams)).Count(&total).Error ; err != nil {
    return total, err
  }
  pR.Log.Info("Get All (Products)")
  return total, nil
}


func (pR *ProductRepository) FilterProducts(searchParams string) func(tx *gorm.DB) *gorm.DB {
  return func(tx *gorm.DB) *gorm.DB {
    return tx.Where("title ILIKE ?", "%" +searchParams + "%").Or("description ILIKE ?", "%" + searchParams + "%")
  }
}

func (pR *ProductRepository) GetProductById(productId string,product *entity.Product) error {
  pR.Log.Infof("Get By Id (Products) : %v", productId)
  return pR.DB.Where("product_id = ? ", productId ).First(&product).Error
}
