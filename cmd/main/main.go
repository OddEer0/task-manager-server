package main

import (
	"github.com/OddEer0/task-manager-server/config"
	"github.com/OddEer0/task-manager-server/internal/presentation/router"
	"log"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	appRouter := router.AppRouter()
	server := http.Server{Addr: cfg.Host, Handler: appRouter}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("failed to start server")
	} else {
		log.Printf("Server started to host: %s", cfg.Host)
	}
}
