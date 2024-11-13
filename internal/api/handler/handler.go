package handler

import (
	"encoding/json"
	"go-project/internal/api/model"
	"go-project/internal/api/services"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	var req model.User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, err := userServices.LoginUser(req.Email, req.Password)
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

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	var req model.User

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newUser, err := userServices.SignupUser(req)

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
