package handler

import (
	"encoding/json"
	"fmt"
	"go-project/internal/api/model"
	userServices "go-project/internal/api/services"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func HandleLogin(service *userServices.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req model.User
		ctx := r.Context()

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		user, accessToken, refreshToken, err := service.LoginUser(ctx, req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res := struct {
			User        model.User
			AccessToken string
		}{
			User:        user,
			AccessToken: accessToken,
		}

		cookie := &http.Cookie{
			Name:  "refreshToken",
			Value: refreshToken,
			// HttpOnly: true,
			// Secure:   true,
			MaxAge:  60 * 60 * 24 * 30,
			Expires: time.Now().Add(time.Hour * 24 * 30),
		}

		http.SetCookie(w, cookie)
		err = json.NewEncoder(w).Encode(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func HandleSignup(service *userServices.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req model.User

		ctx := r.Context()

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		newUser, accessToken, refreshToken, err := service.SignupUser(ctx, req)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res := struct {
			User        model.User
			AccessToken string
		}{
			User:        newUser,
			AccessToken: accessToken,
		}
		cookie := &http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			HttpOnly: true,
			Secure:   true,
			MaxAge:   60 * 60 * 24 * 30,
			Expires:  time.Now().Add(time.Hour * 24 * 30),
		}

		http.SetCookie(w, cookie)
		err = json.NewEncoder(w).Encode(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
}

// func HandleUsers(service *userServices.UserService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		val := r.Context().Value("user")
// 		fmt.Println(val)
// 	}
// }

func HandleRefresh(service *userServices.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cookie, err := r.Cookie("refreshToken")
		if err != nil {
			http.Error(w, " missing", http.StatusNotFound)
		}

		tokenString := cookie.Value
		secretKey := []byte(os.Getenv("REFRESH_TOKEN_SECRET"))

		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "token expired", http.StatusUnauthorized)
			return
		}

		accessToken, err := service.TokenRefresh(ctx, tokenString)
		if err != nil {
			http.Error(w, "error in token creation", http.StatusInternalServerError)
			return
		}

		res := map[string]string{"accessToken": accessToken}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

	}
}
