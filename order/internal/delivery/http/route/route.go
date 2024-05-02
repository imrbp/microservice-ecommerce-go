package route

import (
  delivery "order_service/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
  App *fiber.App
  ProductController *delivery.ProductController 
  CartController *delivery.CartController
  OrderController *delivery.OrderController
}

func (r *RouteConfig) Setup() {
  r.App.Get("/api/products", r.ProductController.GetAll)
  r.App.Get("/api/products/:productId", r.ProductController.GetById)
  r.App.Post("/api/products", r.ProductController.AddToCart)

  r.App.Get("/api/carts", r.CartController.GetAll)
  r.App.Put("/api/carts", r.CartController.AddQuantityProduct)
  r.App.Delete("/api/carts", r.CartController.DeleteProduct)

  r.App.Get("/api/orders", r.OrderController.GetAll)
  r.App.Post("/api/orders", r.OrderController.Checkout)

}
