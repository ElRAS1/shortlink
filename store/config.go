package store

import (
	"sync"
)

type data struct {
	oldlink string
	newlink string
}
type cache struct {
	mu *sync.RWMutex
	ch map[data]struct{}
}

type configDB struct {
	Port     string `env:"PORT" env-default:"5432"`
	Host     string `env:"HOST" env-default:"localhost"`
	Name     string `env:"NAME" env-default:"postgres"`
	User     string `env:"USER" env-default:"user"`
	Password string `env:"PASSWORD"`
}
