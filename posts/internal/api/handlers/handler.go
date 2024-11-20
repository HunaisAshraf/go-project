package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hunaisashraf/go-auth/internal/api/model"
	services "github.com/hunaisashraf/go-auth/internal/api/service"
)

func Health(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is working"))
	}
}

func HandleAddPost(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body model.Posts
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
		var userId interface{} = r.Context().Value("userId")
		body.UserId = string(userId.(string))

		fmt.Println(body)

		post, err := service.CreatePost(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("post created")

		err = json.NewEncoder(w).Encode(post)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func HandleGetPost(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		post, err := service.GetPost(id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode((post))
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
