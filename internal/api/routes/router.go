package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-project/internal/api/handler"
	userServices "go-project/internal/api/services"
)

func Router(service *userServices.UserService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/login", handler.HandleLogin(service))
	r.Post("/signup", handler.HandleSignup(service))

	return r
}
