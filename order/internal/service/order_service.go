package service

import (
	"context"
	delivery "order_service/internal/delivery/messaging"
	"order_service/internal/entity"
	"order_service/internal/model"
	"order_service/internal/model/converter"
	"order_service/internal/repository"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderService struct {
	DB  *gorm.DB
	Log *logrus.Logger

	OrderRepository        *repository.OrderRepository
	CartRepository         *repository.CartRepository
	ProductRepository      *repository.ProductRepository
	OrderProductRepository *repository.OrderProductsRespository

	OrderPublisher *delivery.OrderPublisher
}

func NewOrderService(
	db *gorm.DB,
	log *logrus.Logger,
	orderRepository *repository.OrderRepository,
	cartRepository *repository.CartRepository,
	productRepository *repository.ProductRepository,
	orderProductsRespository *repository.OrderProductsRespository,
	orderPublisher *delivery.OrderPublisher,
) *OrderService {
	return &OrderService{
		DB:                     db,
		Log:                    log,
		OrderRepository:        orderRepository,
		CartRepository:         cartRepository,
		ProductRepository:      productRepository,
		OrderProductRepository: orderProductsRespository,
		OrderPublisher:         orderPublisher,
	}
}

func (oS *OrderService) GetAllOrder(ctx context.Context, orderSearch *model.SearchOder) error {
	var orders []entity.Orders
	if err := oS.OrderRepository.GetByFilter(orderSearch.FilterBy, &orders); err != nil {
		return err
	}

	for _, order := range orders {
		orderSearch.Orders = append(orderSearch.Orders, *converter.OrderToResponse(&order))
	}
	return nil
}

func (oS *OrderService) Checkout(ctx context.Context, orderRequest *model.OrderRequest) error {
	tx := oS.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	orderId := uuid.NewString()

	var products []entity.OrderProducts
	order := entity.Orders{
		OrderId:    orderId,
		Status:     "pending",
		Airwaybill: "",
		Invoice:    "",
		TotalPrice: 0,
	}

	for _, product := range orderRequest.Products {
		var cart entity.Cart
		if err := oS.CartRepository.GetById(product.ProductId, &cart); err != nil {
			return err
		}

		var productR entity.Product
		if err := oS.ProductRepository.GetProductById(product.ProductId, &productR); err != nil {
			return err
		}

		product.Price = int64(product.Quantity) * productR.Price
		products = append(products, *converter.M_OrderProduct_E_OrderProduct(orderId, &product))
		order.TotalPrice = order.TotalPrice + productR.Price
	}

	if err := oS.OrderRepository.CreateOrder(order); err != nil {
		return err
	}

	if err := oS.OrderProductRepository.CreateOrderProducts(&products); err != nil {
		return err
	}

	err := oS.OrderPublisher.Checkout(&order)

	if err != nil {
		return err
	}
	return nil
}
