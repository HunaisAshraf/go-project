package routes

import (
	"github.com/go-chi/chi/v5"
	"go-project/internal/api/handler"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", handler.HandleLogin)
	r.Post("/signup", handler.HandleSignup)

	return r
}
