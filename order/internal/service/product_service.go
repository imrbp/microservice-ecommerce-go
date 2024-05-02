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

type ProductService struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	ProductRepository *repository.ProductRepository
	CartRepository    *repository.CartRepository
}

func NewProductService(
	db *gorm.DB,
	log *logrus.Logger,
	productReposiotry *repository.ProductRepository,
	cartRepository *repository.CartRepository,
) *ProductService {
	return &ProductService{
		DB:                db,
		Log:               log,
		ProductRepository: productReposiotry,
		CartRepository:    cartRepository,
	}
}

func (pS *ProductService) Get(ctx context.Context, searchParams string, request *model.SearchProduct) error {
	var products []entity.Product
	total, err := pS.ProductRepository.GetAllProducts(searchParams, request.Size, request.CurrentPage, &products)
	if err != nil {
		return err
	}
	request.Total = total

	for _, product := range products {
		request.Products = append(request.Products, *converter.ProductToResponse(&product))
	}
	return nil
}

func (pS *ProductService) GetById(ctx context.Context, productId string, productResponse *model.ProductResponse) error {
	var product entity.Product

	if err := pS.ProductRepository.GetProductById(productId, &product); err != nil {
		return err
	}

  *productResponse = *converter.ProductToResponse(&product)
	return nil
}

func (pS *ProductService) AddToCart(ctx context.Context, productCart *model.AddToCartRequest) error {
	tx := pS.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var product entity.Product

	if err := pS.ProductRepository.GetProductById(productCart.ProductId, &product); err != nil {
		return err
	}

	cart := entity.Cart{
    ProductId : product.ProductId,
		Quantity: productCart.Quantity,
	}

	if err := pS.CartRepository.AddToCart(&cart); err != nil {
		return err
	}
	return nil
}
