package router

import (
	"github.com/OddEer0/task-manager-server/internal/presentation/handlers/http"
	"github.com/go-chi/chi/v5"
)

func AppRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/http/v1", func(router chi.Router) {
		router.Post("/auth/registration", http.NewAuthHandler().Registration)
	})

	return router
}
