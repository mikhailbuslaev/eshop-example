package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mikhailbuslaev/eshop/storage"
)

type Server struct {
	Storage *storage.Storage
	Router  *gin.Engine
}

func New() *Server {
	s := &Server{}
	s.Storage = storage.New()
	s.Router = gin.Default()
	// cards handlers
	s.Router.POST("api/cards", s.getCardsHandler)
	s.Router.POST("api/cards_count", s.getCardsQuantityHandler)
	s.Router.POST("api/card", s.getCardHandler)
	s.Router.POST("api/cards/delete", s.deleteCardHandler)
	s.Router.POST("api/cards/add", s.addCardHandler)
	s.Router.POST("api/cards/update", s.updateCardHandler)
	// shopping carts handlers
	s.Router.POST("api/shopping_cart", s.getShoppingCartHandler)
	s.Router.POST("api/shopping_cart/clear", s.clearShoppingCartHandler)
	s.Router.POST("api/shopping_cart/add_item", s.addToShoppingCartHandler)
	s.Router.POST("api/shopping_cart/delete_item", s.deleteFromShoppingCartHandler)
	// authorization handlers
	s.Router.POST("api/get_token", s.getTokenHandler)
	// orders handlers
	s.Router.POST("api/orders", s.getOrdersHandler)
	s.Router.POST("api/orders_count", s.getOrdersQuantityHandler)
	s.Router.POST("api/order", s.getOrderHandler)
	s.Router.POST("api/orders/delete", s.deleteOrderHandler)
	s.Router.POST("api/orders/add", s.addOrderHandler)
	return s
}

func (s *Server) Run() {
	s.Router.Run(":1111")
}
