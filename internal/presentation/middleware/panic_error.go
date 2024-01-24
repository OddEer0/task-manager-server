package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

func PanicMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error(
						"err", err,
						"trace", debug.Stack(),
					)
				}
			}()
		})
	}
}
