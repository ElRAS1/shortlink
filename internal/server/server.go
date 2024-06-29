package server

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	ctx, cancel := context.WithCancel(context.Background())
	server := &server{}
	server.srv = &http.Server{Addr: cfg.Port, Handler: server.Routes()}
	server.ConfigureLogger(cfg.Loglevel, cfg.Configlog)
	server.gracefulShutdown(cancel)
	slog.SetDefault(server.logger)
	log.Printf("app starting in port%s\n", server.srv.Addr)

	err = server.ConfigureStore(ctx)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connecting in database")

	return server.srv
}

func (s *server) ConfigureStore(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		s.store = store.New()
		if err := s.store.Open(); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) gracefulShutdown(cancel context.CancelFunc) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-sig:
				err := s.store.Db.Close()
				if err != nil {
					log.Println(err)
				}
				cancel()
				log.Println("stopping server in graceful shutdown...")
				return
			default:
				continue
			}
		}
	}()
}
