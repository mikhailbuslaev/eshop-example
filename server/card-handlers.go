package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikhailbuslaev/eshop/model"
)

func (s *Server) getCardsHandler(ctx *gin.Context) {
	quantity, err := strconv.Atoi(ctx.PostForm("quantity"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if quantity <= 0 || quantity > 100 {
		quantity = 20
	}
	cards, err := s.Storage.GetCards(quantity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": *cards})
}

func (s *Server) getCardHandler(ctx *gin.Context) {
	cardId := ctx.PostForm("id")
	card, err := s.Storage.GetCard(cardId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": *card})
}

func (s *Server) deleteCardHandler(ctx *gin.Context) {
	cardId := ctx.PostForm("id")
	err := s.Storage.DeleteCard(cardId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "card successfully removed"})
}

func (s *Server) addCardHandler(ctx *gin.Context) {
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

	c.Id = generateId(10)
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

	if err = ctx.SaveUploadedFile(file, "website/"+c.PicturePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "card successfully added"})
}

func (s *Server) updateCardHandler(ctx *gin.Context) {
	cardId := ctx.Param("id")
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
	count, err := strconv.Atoi(ctx.PostForm("count"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if count != 0 {
		card.Count = count
	}

	file, err := ctx.FormFile("picture")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if file.Size == 0 {
		if err = ctx.SaveUploadedFile(file, card.PicturePath); err != nil {
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
