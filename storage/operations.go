package storage

import (
	"database/sql"
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

func (s *Storage) GetUrlDb(oldlink string) (data, error) {
	sqlStatement := `SELECT * FROM link WHERE oldlink = ? RETURNING id, oldlink, newlink`
	rows := s.Db.QueryRowx(sqlStatement, oldlink)
	result := data{}
	err := rows.StructScan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return data{}, fmt.Errorf("no record found for oldlink: %s", oldlink)
		}
		return data{}, err
	}
	return result, nil
}
