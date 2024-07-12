package storage

import (
	"fmt"
	"sync"

	"github.com/ELRAS1/shortlink/internal/logger"
)

type Cache struct {
	mu *sync.RWMutex
	ch map[string]data
}

func (s *Storage) CachedUrl() {
	s.Cache.mu.Lock()
	s.Cache.ch[s.Data.newlink] = *s.Data
	logger.Logger.Debug("Data saved successfully", "data", fmt.Sprintf("%+v", s.Data))
	s.Cache.mu.Unlock()
}
