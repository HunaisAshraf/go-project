package handler

import (
	"encoding/json"
	"go-project/internal/api/model"
	"go-project/internal/api/services"
	"net/http"
	"time"
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
