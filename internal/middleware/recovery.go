package middleware

import (
	"log/slog"
	"net/http"
)

func Recovery(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("recovery middleware found a panic: %v", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
