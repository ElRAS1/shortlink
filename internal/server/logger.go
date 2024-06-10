package server

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)

func (s *server) ConfigureLogger(level int, cfg string) {
	switch cfg {
	case "dev":
		s.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(level)}))
	case "prod":
		s.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(level)}))
	}
}

func (s *server) logMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Запрос:", r.URL)
		next.ServeHTTP(w, r)
	})
}
