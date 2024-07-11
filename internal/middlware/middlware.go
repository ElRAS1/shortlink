package middlware

import (
	"log"
	"net/http"
)

func Middlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Запрос:", r.URL)
		next.ServeHTTP(w, r)
	})
}
