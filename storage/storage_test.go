package storage

import (
	"testing"

	"github.com/ELRAS1/shortlink/internal/logger"
)

// type StoragerTest interface {
// 	urlToSave()
// 	CachedUrl()
// }

func TestSaveUrl(t *testing.T) {
	logger.ConfigureLogger(0, "dev")
	testStorage := New()
	testStorage.Data.oldlink = "https:/yandex.ru"
	testStorage.Data.newlink = "https:/ya.ru"
	testStorage.Data.id = 0

	testStorage.CachedUrl()

	data, ok := testStorage.Cache.ch["https:/yandex.ru"]
	if !ok || data.newlink != "https:/ya.ru" {
		t.Errorf("expected data[oldlink] = https:/ya.ru, res = %v", data.newlink)
	}
}
