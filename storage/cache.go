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
	s.Cache.ch[s.Data.oldlink] = *s.Data
	logger.Logger.Debug("Data saved successfully", "data", fmt.Sprintf("%+v", s.Data))
	s.Cache.mu.Unlock()
}

func (s *Storage) GetCache(link string) (data, bool) {
	s.Cache.mu.RLock()
	data, ok := s.Cache.ch[link]
	if !ok {
		logger.Logger.Debug("no data found in the cache")
		return data, false
	}
	logger.Logger.Debug("the data is taken from the cache")
	s.Cache.mu.RUnlock()
	return data, true
}
