package middleware

import (
	//"errors"
	"context"
	"fmt"
	"learn_gqlgen/auth/services"
	"net/http"
	"strings"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := VerifyHeaderToken(w, r)
		ctx := r.Context()
		if err != nil {
			// Insert error in context ( used in resolver )
			ctx = context.WithValue(ctx, "authError", err)
		} else {
			ctx = context.WithValue(ctx, "userID", userID)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func VerifyHeaderToken(w http.ResponseWriter, r *http.Request) (string, error) {
	// Verify if authorizatioin exists in Header
	verification := r.Header.Get("Authorization")
	if verification == "" {
		return "", fmt.Errorf("Unauthorized")
	}

	// Transform authorization in array
	partsAuthorization := strings.Split(verification, " ")
	if len(partsAuthorization) != 2 {
		return "", fmt.Errorf("Unauthorized")
	}
	// Verify if first part is write "Bearer"
	if partsAuthorization[0] != "Bearer" {
		return "", fmt.Errorf("Unauthorized")
	}

	token := partsAuthorization[1]
	userID, err := services.DecodeToken(token)
	if err != nil {
		return "", err
	}
	return userID, nil
}