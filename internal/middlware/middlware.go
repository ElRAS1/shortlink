package middlware

import (
	"fmt"
	"net/http"

	"github.com/ELRAS1/shortlink/internal/logger"
)

func Middlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Logger.Info(fmt.Sprintf("Запрос: %s", r.URL))
		next.ServeHTTP(w, r)
	})
}
