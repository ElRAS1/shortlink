package server

import (
	"net/http"

	"github.com/ELRAS1/shortlink/internal/middlware"
)

func (s *server) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/shortlink", middlware.Middlware(http.HandlerFunc(s.handleshort)))
	return mux
}

func (s *server) handleshort(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello world"))

	if err != nil {
		s.logger.Error(err.Error())
	}
}
