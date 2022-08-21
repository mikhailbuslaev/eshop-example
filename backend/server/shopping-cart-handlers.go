package server

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/mikhailbuslaev/eshop/model"
)

func CartToJson(cart *model.ShoppingCart) error {
	jsonMap := make(map[string][]model.CartItem)
	jsonMap["items"] = cart.Items
	buf, err := json.Marshal(jsonMap)
	if err != nil {
		return err
	}
	cart.JsonItems = buf
	return nil
}

func (s *Server) getShoppingCartHandler(ctx *gin.Context) {
	token := ctx.PostForm("token")
	cartId, isValid := validateToken(token)
	if !isValid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid authorization token"})
		return
	}

	cart, err := s.Storage.GetShoppingCart(cartId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{"message": *cart})
}

func (s *Server) clearShoppingCartHandler(ctx *gin.Context) {
	token := ctx.PostForm("token")
	cartId, isValid := validateToken(token)
	if !isValid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid authorization token"})
		return
	}

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
	token := ctx.PostForm("token")
	cartId, isValid := validateToken(token)
	if !isValid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid authorization token"})
		return
	}

	cardId := ctx.PostForm("cardid")
	if cardId == "" {
		ctx.JSON(http.StatusOK, gin.H{"error": "cardId is not defined"})
		return
	}

	stringCount := ctx.PostForm("count")
	if stringCount == "" {
		ctx.JSON(http.StatusOK, gin.H{"error": "count is not defined"})
		return
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
		return
	}
	cart.Items = append(cart.Items, item)

	if err = s.Storage.ClearShoppingCart(cartId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = CartToJson(cart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = s.Storage.UpdateShoppingCart(cart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{"message": "item successfully added"})
}

func (s *Server) deleteFromShoppingCartHandler(ctx *gin.Context) {
	isExist := false
	
	token := ctx.PostForm("token")
	cartId, isValid := validateToken(token)
	if !isValid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid authorization token"})
		return
	}

	cartItemId := ctx.PostForm("cartitemid")
	if cartItemId == "" {
		ctx.JSON(http.StatusOK, gin.H{"error": "cartItemId is not defined"})
		return
	}

	cart, err := s.Storage.GetShoppingCart(cartId)
	if err != nil {
		fmt.Println("error in get cart")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	length := len(cart.Items)
	for i := 0; i < length; i++ {
		if cart.Items[i].Id == cartItemId {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			isExist = true
			break
		}
	}

	if !isExist {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "there is no cart item with this id"})
		return
	}

	if err = s.Storage.ClearShoppingCart(cartId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = CartToJson(cart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = s.Storage.UpdateShoppingCart(cart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{"message": "item successfully deleted"})
}