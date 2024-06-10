package server

import (
	"log"
	"log/slog"
	"net/http"

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
	// store
}

func NewServer(cfg config) *server {
	return &server{}
}

func StartServer() *http.Server {
	cfg := config{}
	err := cleanenv.ReadConfig("config/config.yaml", &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	server := NewServer(cfg)
	server.srv = &http.Server{Addr: cfg.Port, Handler: server.Routes()}
	server.ConfigureLogger(cfg.Loglevel, cfg.Configlog)
	slog.SetDefault(server.logger)
	log.Printf("app starting in port%s\n", server.srv.Addr)
	return server.srv
}
