package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type CustomJWTClaims struct {
	UserId string `json:"id"`
	jwt.StandardClaims
}

const jwtKey = "key_example" // in future will be in config, git ignored

func createToken(id string) (string, error){
	claims := CustomJWTClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Issuer:    "eshop_test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return stringToken, nil
}


func (s *Server) getTokenHandler(ctx *gin.Context) {
	id := ctx.PostForm("id")
	if id == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "id undefined"})
		return
	}

	password := ctx.PostForm("password")
	if password == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "password undefined"})
		return
	}
	realPass, err := s.Storage.GetPassword(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if realPass != password {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid password"})
		return
	}

	stringToken, err := createToken(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "token creation failed"})
		return
	}
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{"message": stringToken})
}