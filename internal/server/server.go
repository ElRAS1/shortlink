package server

import (
	"context"
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
	Srv    *http.Server
	logger *slog.Logger
	store  *store.Store
}

func StartServer(ctx context.Context) server {
	cfg := config{}
	err := cleanenv.ReadConfig("config/config.yaml", &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	server := &server{}
	server.Srv = &http.Server{Addr: cfg.Port, Handler: server.Routes()}
	server.ConfigureLogger(cfg.Loglevel, cfg.Configlog)
	slog.SetDefault(server.logger)
	log.Printf("app starting in port%s\n", server.Srv.Addr)
	err = server.ConfigureStore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connecting in database")

	return *server
}

func (s *server) ConfigureStore(ctx context.Context) error {
	s.store = store.New()
	if err := s.store.Open(ctx); err != nil {
		return err
	}
	return nil
}

func Finish(ctx context.Context, s server) error {
	if err := s.store.Db.Close(); err != nil {
		s.logger.Error(err.Error())
	}
	if err := s.Srv.Close(); err != nil {
		s.logger.Error(err.Error())
	}
	log.Println("stopping server in graceful shutdown...")
	select {
	case <-ctx.Done():
		log.Println("context cancelled")
	default:
	}
	return nil
}
