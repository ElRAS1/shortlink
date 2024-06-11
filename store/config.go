package store

import (
	"database/sql"
	"sync"
)

type database struct {
	db    *sql.DB
	DbUrl string
}
type data struct {
	oldlink string
	newlink string
}
type cache struct {
	mu *sync.RWMutex
	ch map[data]struct{}
}

type configDB struct {
	port     int
	host     string
	user     string
	password string
	dbname   string
}
