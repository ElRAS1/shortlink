package server

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"github.com/ELRAS1/shortlink/internal/logger"
	"github.com/ELRAS1/shortlink/storage"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Loglevel  int    `yaml:"level"`
	Port      string `yaml:"addr" env-default:"8080"`
	Configlog string `yaml:"configlogger"`
}

type server struct {
	Srv   *http.Server
	store *storage.Storage
}

func StartServer(ctx context.Context) server {
	cfg := config{}
	err := cleanenv.ReadConfig("config/config.yaml", &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	server := &server{}
	server.Srv = &http.Server{Addr: cfg.Port, Handler: server.Routes()}
	// server.logger = logger.ConfigureLogger(cfg.Loglevel, cfg.Configlog)
	logger.ConfigureLogger(cfg.Loglevel, cfg.Configlog)
	slog.SetDefault(logger.Logger)
	log.Printf("app starting in port%s\n", server.Srv.Addr)
	err = server.ConfigureStore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	logger.Logger.Info("connecting in database")
	return *server
}

func (s *server) ConfigureStore(ctx context.Context) error {
	s.store = storage.New()
	if err := storage.Open(ctx, s.store); err != nil {
		return err
	}
	return nil
}

func Finish(ctx context.Context, s server) error {
	if err := s.store.Db.Close(); err != nil {
		logger.Logger.Error(err.Error())
	}
	log.Println("stopping server in graceful shutdown...")
	select {
	case <-ctx.Done():
		logger.Logger.Info("context cancelled")
		return nil
	default:
	}
	return nil
}
