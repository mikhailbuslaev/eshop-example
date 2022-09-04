package server

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func validateToken(tokenString string) (string, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return "", false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims.Valid() == nil && claims["id"] != nil {
			return claims["id"].(string), true
		}
	}
	return "", false
}

func (s *Server) getShoppingCartCost(cartId string) (float64, error) {
	var summaryCost float64

	cart, err := s.Storage.GetShoppingCart(cartId)
	if err != nil {
		return 0.0, err
	}

	for i := range cart.Items {
		card, err := s.Storage.GetCard(cart.Items[i].CardId)
		if err != nil {
			return 0.0, err
		}
		cardCost, err := strconv.ParseFloat(card.Price, 64)
		if err != nil {
			return 0.0, err
		}
		summaryCost += cardCost * float64(cart.Items[i].Count)
	}
	return summaryCost, nil
}
