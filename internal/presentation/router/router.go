package router

import (
	"net/http"

	appErrors "github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	"github.com/OddEer0/task-manager-server/internal/infrastructure/logger"
	"github.com/OddEer0/task-manager-server/internal/presentation/handlers"
	customMiddleware "github.com/OddEer0/task-manager-server/internal/presentation/middleware"
	"github.com/go-chi/chi/v5"
)

func AppRouter(router *chi.Mux) {
	appHandler := handlers.NewAppHandler()

	middleware := appErrors.LoggingMiddleware(logger.NewLogger())

	router.Route("/http/v1", func(router chi.Router) {
		router.Post("/auth/registration", middleware(appHandler.AuthHandler.Registration))
		router.Post("/auth/login", middleware(appHandler.AuthHandler.Login))
		router.Post("/auth/logout", middleware(appHandler.AuthHandler.Logout))
		router.Get("/auth/refresh", middleware(appHandler.AuthHandler.Refresh))

		router.Route("/", func(router chi.Router) {
			router.Use(customMiddleware.AuthMiddleware)
			router.Get("/hello", middleware(func(res http.ResponseWriter, req *http.Request) error {
				res.Write([]byte("Hello auth!"))
				return nil
			}))
		})
	})
}
