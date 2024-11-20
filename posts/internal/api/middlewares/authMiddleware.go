package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		tokenString := strings.TrimPrefix(header, "Bearer ")
		secretKey := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))

		fmt.Println(tokenString)
		fmt.Println(os.Getenv("ACCESS_TOKEN_SECRET"))

		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Println("error signing method")
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		fmt.Println(token)
		if err != nil || !token.Valid {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		fmt.Println("tonken valid")
		if claims, ok := token.Claims.(*jwt.MapClaims); ok {
			userId := (*claims)["_id"].(string)

			ctx := context.WithValue(r.Context(), "userId", userId)

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

	})
}
