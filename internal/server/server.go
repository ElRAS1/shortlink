package server

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/ELRAS1/shortlink/store"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Loglevel  int    `yaml:"level"`
	Port      string `yaml:"addr" env-default:"8080"`
	Configlog string `yaml:"configlogger"`
}

type server struct {
	srv    *http.Server
	logger *slog.Logger
	store  *store.Store
}

func StartServer() *http.Server {
	cfg := config{}
	err := cleanenv.ReadConfig("config/config.yaml", &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	server := &server{}
	server.srv = &http.Server{Addr: cfg.Port, Handler: server.Routes()}
	server.ConfigureLogger(cfg.Loglevel, cfg.Configlog)
	slog.SetDefault(server.logger)
	log.Printf("app starting in port%s\n", server.srv.Addr)

	err = server.ConfigureStore()

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connecting in database")

	return server.srv
}

func (s *server) ConfigureStore() error {
	s.store = store.New()

	if err := s.store.Open(); err != nil {
		return err
	}

	return nil
}
