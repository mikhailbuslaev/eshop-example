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
	s.Router.GET("/cards", s.getCardsHandler)
	s.Router.GET("/cards/:card", s.getCardHandler)
	s.Router.POST("/cards/remove", s.removeCardHandler)
	s.Router.POST("/cards/add", s.putCardHandler)
	return s
}

func (s *Server) Run() {
	s.Router.Run(":1111")
}
