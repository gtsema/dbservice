package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDB(url string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", url)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
