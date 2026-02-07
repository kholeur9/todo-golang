package services

import (
	//"fmt"
	"fmt"
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

func DecodeToken(tokenAuth string) (string, error) {
	token, err := jwt.Parse(tokenAuth, func(tokenAuth *jwt.Token) (any, error) {
		return key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return "", fmt.Errorf("Unauthorized")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Unauthorized")
	}
	userID := claims["sub"].(string)
	return userID, nil
}