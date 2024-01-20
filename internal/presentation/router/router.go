package router

import (
	"github.com/OddEer0/task-manager-server/internal/presentation/handlers"
	"github.com/go-chi/chi/v5"
)

func AppRouter(router *chi.Mux) {
	appHandler := handlers.NewAppHandler()

	router.Route("/http/v1", func(router chi.Router) {
		router.Post("/auth/registration", appHandler.Registration)
	})
}
