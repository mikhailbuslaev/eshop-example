package server

import (
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mikhailbuslaev/eshop/model"
)

func (s *Server) getOrdersHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	token := ctx.PostForm("token")
	userId, isValid := validateToken(token)
	if !isValid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid authorization token"})
		return
	}

	quantity, err := strconv.Atoi(ctx.PostForm("quantity"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if quantity <= 0 || quantity > 100 {
		quantity = 20
	}

	startRow, err := strconv.Atoi(ctx.PostForm("startrow"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	orders, err := s.Storage.GetOrders(userId, quantity, startRow)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": *orders})
}

func (s *Server) getOrderHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	token := ctx.PostForm("token")
	userId, isValid := validateToken(token)
	if !isValid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid authorization token"})
		return
	}

	orderId := ctx.PostForm("id")
	order, err := s.Storage.GetOrder(userId, orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": *order})
}

func (s *Server) deleteOrderHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	token := ctx.PostForm("token")
	userId, isValid := validateToken(token)
	if !isValid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid authorization token"})
		return
	}

	orderId := ctx.PostForm("id")
	err := s.Storage.DeleteOrder(userId, orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "order successfully removed"})
}

func (s *Server) addOrderHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	o := &model.Order{}
	var err error
	o.Delivery.Address = ctx.PostForm("delivery-address")
	o.Delivery.Zip = ctx.PostForm("delivery-zip")
	o.Delivery.Receiver = ctx.PostForm("delivery-receiver")
	token := ctx.PostForm("token")
	shoppingCartId, isValid := validateToken(token)
	if !isValid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid authorization token"})
		return
	}

	cart, err := s.Storage.GetShoppingCart(shoppingCartId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	o.Items = cart.Items
	o.Cost, err = s.getShoppingCartCost(shoppingCartId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	o.Id = randomString(20)
	o.UserId = shoppingCartId
	jsonData, err := json.Marshal(*o)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	o.JsonData = string(jsonData)

	if err = s.Storage.AddOrder(o); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "order successfully added"})
}

func (s *Server) getOrdersQuantityHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	token := ctx.PostForm("token")
	userId, isValid := validateToken(token)
	if !isValid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid authorization token"})
		return
	}

	quantity, err := s.Storage.GetOrdersQuantity(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": quantity})
}