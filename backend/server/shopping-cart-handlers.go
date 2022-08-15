package server

import (
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/mikhailbuslaev/eshop/model"
)

func CartToJson(cart *model.ShoppingCart) error {
	buf, err := json.Marshal(cart.Items)
	if err != nil {
		return err
	}
	cart.JsonItems = buf
	return nil
}

func (s *Server) getShoppingCartHandler(ctx *gin.Context) {
	cartId := ctx.PostForm("id")
	cart, err := s.Storage.GetShoppingCart(cartId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{"message": cart.Items})
}

func (s *Server) clearShoppingCartHandler(ctx *gin.Context) {
	cartId := ctx.PostForm("id")
	err := s.Storage.ClearShoppingCart(cartId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{"message": "card successfully cleared"})
}

func (s *Server) addToShoppingCartHandler(ctx *gin.Context) {
	item := model.CartItem{}

	cartId := ctx.PostForm("cartid")
	if cartId == "" {
		ctx.JSON(http.StatusOK, gin.H{"error": "cartId is not defined"})
	}

	cardId := ctx.PostForm("cardid")
	if cardId == "" {
		ctx.JSON(http.StatusOK, gin.H{"error": "cardId is not defined"})
	}

	stringCount := ctx.PostForm("count")
	if stringCount == "" {
		ctx.JSON(http.StatusOK, gin.H{"error": "count is not defined"})
	}

	count, err := strconv.Atoi(stringCount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	item.Id = randomString(20)
	item.CardId = cardId
	item.Count = count

	cart, err := s.Storage.GetShoppingCart(cartId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	cart.Items = append(cart.Items, item)

	if err = s.Storage.ClearShoppingCart(cartId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err = CartToJson(cart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if err = s.Storage.UpdateShoppingCart(cart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{"message": "item successfully added"})
}

func (s *Server) deleteFromShoppingCartHandler(ctx *gin.Context) {
	cartItemId := ctx.PostForm("cartitemid")
	cartId := ctx.PostForm("id")

	cart, err := s.Storage.GetShoppingCart(cartId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	for i := range cart.Items {
		if cart.Items[i].Id == cartItemId {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
		}
	}

	if err = s.Storage.ClearShoppingCart(cartId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err = CartToJson(cart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err = s.Storage.UpdateShoppingCart(cart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{"message": "item successfully deleted"})
}