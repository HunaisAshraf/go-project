package handler

import (
	"encoding/json"
	"go-project/internal/api/model"
	"go-project/internal/api/services"
	"net/http"
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

		user, err := service.LoginUser(ctx, req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(user)

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

		newUser, err := service.SignupUser(ctx, req)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(newUser)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
}
