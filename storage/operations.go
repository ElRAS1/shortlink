package storage

import (
	"fmt"
)

func (s *Storage) SaveUrl(urlToSave string, alias string) error {
	if urlToSave == "" || alias == "" {
		return fmt.Errorf("url and alias are required")
	}
	sqlStatement := `INSERT INTO urls (url, alias) VALUES (?, ?) RETURNING id`
	row := s.Db.QueryRow(sqlStatement, urlToSave, alias)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	s.Data.id = uint8(id)
	s.Data.oldlink = urlToSave
	s.Data.newlink = alias
	s.CachedUrl()
	return nil
}
