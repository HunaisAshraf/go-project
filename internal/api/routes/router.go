package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-project/internal/api/handler"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/login", handler.HandleLogin)
	r.Post("/signup", handler.HandleSignup)

	return r
}
