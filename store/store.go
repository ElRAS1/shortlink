package store

import (
	"context"
	"fmt"
	"os"
	"time"

	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Store struct {
	Db    *sqlx.DB
	Cache *cache
	Data  *data
}

func New() *Store {
	chc := cache{mu: &sync.RWMutex{}, ch: make(map[data]struct{})}
	data := data{oldlink: "", newlink: ""}
	return &Store{Db: &sqlx.DB{}, Cache: &chc, Data: &data}
}

func (s *Store) Open(ctx context.Context) error {
	cfg := configDB{}
	err := s.loadEnv(&cfg)
	if err != nil {
		return err
	}
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	s.Db, err = sqlx.ConnectContext(ctx, "postgres", dbUrl)
	if err != nil {
		return err
	}
	if err = s.Db.PingContext(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Store) loadEnv(cfg *configDB) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	cfg.Host = os.Getenv("DB_HOST")
	cfg.Name = os.Getenv("DB_NAME")
	cfg.Password = os.Getenv("DB_PASSWORD")
	cfg.Port = os.Getenv("DB_PORT")
	cfg.User = os.Getenv("DB_USER")

	return nil
}
