package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/hunaisashraf/go-auth/internal/api/handlers"
	"github.com/hunaisashraf/go-auth/internal/api/middlewares"
	services "github.com/hunaisashraf/go-auth/internal/api/service"
)

func Router(service *services.Service) *chi.Mux {

	r := chi.NewRouter()

	// r.Route("/api/posts", func(r chi.Router) {

	r.Get("/health", handlers.Health(service))
	r.With(middlewares.VerifyToken).Post("/post", handlers.HandleAddPost(service))
	r.With(middlewares.VerifyToken).Get("/post/{id}", handlers.HandleGetPost(service))
	// })

	return r

}
