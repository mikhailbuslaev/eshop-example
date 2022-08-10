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
	s.Router.POST("api/card", s.getCardHandler)
	s.Router.POST("api/cards/delete", s.deleteCardHandler)
	s.Router.POST("api/cards/add", s.addCardHandler)
	s.Router.POST("api/cards/update", s.updateCardHandler)
	// catalog handler
//	s.Router.Static("/catalog", "./website")
	return s
}

func (s *Server) Run() {
	s.Router.Run(":1111")
}
