package main

import (
	"log"
	"net/http"

	"github.com/OddEer0/task-manager-server/config"
	loggerLib "github.com/OddEer0/task-manager-server/internal/infrastructure/logger"
	"github.com/OddEer0/task-manager-server/internal/presentation/middleware"
	appRouter "github.com/OddEer0/task-manager-server/internal/presentation/router"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("main load config error: %v", err)
	}
	logger := loggerLib.SetupLogger(cfg.Env)

	router := chi.NewRouter()

	router.Use(middleware.PanicMiddleware(logger))
	router.Use(loggerLib.LoggingMiddleware(logger))

	appRouter.AppRouter(router)

	server := http.Server{Addr: cfg.Host, Handler: router}

	logger.Info("Server started on host", "host", cfg.Host)

	if err := server.ListenAndServe(); err != nil {
		logger.Error("Failed to start server")
	}
}
