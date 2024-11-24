package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func AccessLog(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		entryTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(entryTime)
		slog.Info("access log",
			slog.String("method", r.Method),
			slog.String("uri", r.RequestURI),
			slog.Duration("duration", duration),
		)
	})
}
