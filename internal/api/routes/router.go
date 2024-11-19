package routes

import (
	"go-project/internal/api/handler"
	userServices "go-project/internal/api/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(service *userServices.UserService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", handler.HandleLogin(service))
		r.Post("/signup", handler.HandleSignup(service))
		r.Get("/refresh-token", handler.HandleRefresh(service))
		// r.With(middlewares.AuthMiddleware).Get("/users", handler.HandleUsers(service))
	})

	return r
}
