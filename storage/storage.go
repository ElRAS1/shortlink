package storage

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

type Storage struct {
	Db    *sqlx.DB
	Cache *Cache
	Data  *data
}

type Storager interface {
	SaveUrl(urlToSave string, alias string) error
	CachedUrl()
}

func New() *Storage {
	chc := Cache{mu: &sync.RWMutex{}, ch: make(map[string]data)}
	data := data{oldlink: "", newlink: "", id: 0}
	return &Storage{Db: &sqlx.DB{}, Cache: &chc, Data: &data}
}

func Open(ctx context.Context, s *Storage) error {
	cfg := configDB{}
	err := loadEnv(&cfg)
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

func loadEnv(cfg *configDB) error {
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
