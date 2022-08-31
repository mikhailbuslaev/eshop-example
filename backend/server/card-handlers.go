package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikhailbuslaev/eshop/model"
)

func (s *Server) getCardsHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
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
	
	cards, err := s.Storage.GetCards(quantity, startRow)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": *cards})
}

func (s *Server) getCardHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	cardId := ctx.PostForm("id")
	card, err := s.Storage.GetCard(cardId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": *card})
}

func (s *Server) deleteCardHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	cardId := ctx.PostForm("id")
	err := s.Storage.DeleteCard(cardId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "card successfully removed"})
}

func (s *Server) addCardHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	c := &model.Card{}
	var err error
	c.Title = ctx.PostForm("title")
	c.Price = ctx.PostForm("price")
	c.Description = ctx.PostForm("description")
	c.Count, err = strconv.Atoi(ctx.PostForm("count"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Id = randomString(10)
	c.PicturePath = "img/" + c.Id + ".jpg"
	file, err := ctx.FormFile("picture")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = s.Storage.AddCard(c); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = ctx.SaveUploadedFile(file, "../frontend/public/"+c.PicturePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "card successfully added"})
}

func (s *Server) updateCardHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	cardId := ctx.PostForm("id")
	println(cardId)
	card, err := s.Storage.GetCard(cardId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if title := ctx.PostForm("title"); title != "" {
		card.Title = title
	}
	if price := ctx.PostForm("price"); price != "" {
		card.Price = price
	}
	if description := ctx.PostForm("description"); description != "" {
		card.Description = description
	}
	stringCount := ctx.PostForm("count")
	if stringCount != "" {
		count, err := strconv.Atoi(stringCount)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		card.Count = count
	}

	file, err := ctx.FormFile("picture")
	if err != nil && err != http.ErrMissingFile {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err != http.ErrMissingFile {
		if err = ctx.SaveUploadedFile(file, "../frontend/public/" + card.PicturePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	if err := s.Storage.UpdateCard(card); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "card successfully updated"})
}

func (s *Server) getCardsQuantityHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	quantity, err := s.Storage.GetCardsQuantity()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": quantity})
}
