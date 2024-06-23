package store

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Store struct {
	Db    *sql.DB
	Cache *cache
	Data  *data
}

func New() *Store {
	chc := cache{mu: &sync.RWMutex{}, ch: make(map[data]struct{})}
	data := data{oldlink: "", newlink: ""}
	return &Store{Db: &sql.DB{}, Cache: &chc, Data: &data}
}

func (s *Store) Open() error {
	cfg := configDB{}
	err := s.LoadEnv(&cfg)
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(cfg.Port)
	if err != nil {
		return err
	}
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, port, cfg.Name)

	s.Db, err = sql.Open("postgres", dbUrl)
	defer s.Db.Close()

	if err != nil {
		return err
	}

	if err = s.Db.Ping(); err != nil {
		return err
	}
	return nil
}

func (s *Store) LoadEnv(cfg *configDB) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	cfg.Host = os.Getenv("HOST")
	cfg.Name = os.Getenv("DB_NAME")
	cfg.Password = os.Getenv("PASSWORD")
	cfg.Port = os.Getenv("PORT")
	cfg.User = os.Getenv("USER")

	return nil
}
