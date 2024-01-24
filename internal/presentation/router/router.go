package router

import (
	appErrors "github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	"github.com/OddEer0/task-manager-server/internal/infrastructure/logger"
	"github.com/OddEer0/task-manager-server/internal/presentation/handlers"
	"github.com/go-chi/chi/v5"
)

func AppRouter(router *chi.Mux) {
	appHandler := handlers.NewAppHandler()

	middleware := appErrors.LoggingMiddleware(logger.NewLogger())

	router.Route("/http/v1", func(router chi.Router) {
		router.Post("/auth/registration", middleware(appHandler.Registration))
		router.Post("/auth/login", middleware(appHandler.Login))
	})
}
