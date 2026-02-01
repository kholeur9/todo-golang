package services

import (
	//"fmt"
	"learn_gqlgen/auth/entity"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte(os.Getenv("JWT_SECRET"))

func GenerateAccessToken(user entity.User) string {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"iat": time.Now(),
		"exp": time.Now().Add(time.Minute*30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Println("JWT Token failed:", err)
		return ""
	}
	return tokenString
}