package server

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/mikhailbuslaev/eshop/card"

	"github.com/gin-gonic/gin"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateCardId() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (s *Server) getCardsHandler(ctx *gin.Context) {
	cards, err := s.Storage.GetCards()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": *cards})
}

func (s *Server) getCardHandler(ctx *gin.Context) {
	cardId := ctx.Param("card")
	card, err := s.Storage.GetCard(cardId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": *card})
}

func (s *Server) removeCardHandler(ctx *gin.Context) {
	cardId := ctx.PostForm("id")
	err := s.Storage.RemoveCard(cardId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "card successfully removed"})
}

func (s *Server) putCardHandler(ctx *gin.Context) {
	c := &card.Card{}
	var err error
	c.Title = ctx.PostForm("title")
	c.Price = ctx.PostForm("price")
	c.Description = ctx.PostForm("description")
	c.Count, err = strconv.Atoi(ctx.PostForm("count"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Id = generateCardId()
	c.PicturePath = "static/pictures/" + c.Id
	file, err := ctx.FormFile("picture")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = s.Storage.PutCard(c); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = ctx.SaveUploadedFile(file, c.PicturePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "card successfully added"})
}
