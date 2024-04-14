package db

import "database/sql"

type SQLiteDB struct {
	*sql.DB
}

func NewSQLiteDB(dataSourceName string) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &SQLiteDB{db}, nil
}

func (db *SQLiteDB) Close() error {
	return db.DB.Close()
}

func (db *SQLiteDB) Initialize() error {
	const createTableSQL = `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        chatId INTEGER NOT NULL,
        name TEXT NOT NULL,
        hash TEXT NOT NULL,
        salt TEXT NOT NULL
    );`

	_, err := db.Exec(createTableSQL)
	return err
}
