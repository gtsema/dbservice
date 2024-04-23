package repository

import (
	"database/sql"
	"dbservice/internal/models/db"
	"fmt"
)

const usersTable = "users"

type SqliteUserRepository struct {
	db *sql.DB
}

func newSqliteUserRepository(database *sql.DB) *SqliteUserRepository {
	return &SqliteUserRepository{db: database}
}

func (r SqliteUserRepository) CreateUser(user db.User) (db.User, error) {
	var user_ db.User
	query := fmt.Sprintf("INSERT INTO %s (chatId, name, hash, salt) VALUES ('$1', '$2', '$3', '$4');", usersTable)

	// SELECT (id, chatId, name, hash, salt) FROM $0 WHERE rowid = last_insert_rowid();

	row := r.db.QueryRow(query, user.ChatId, user.Name, user.Hash, user.Salt)
	if err := row.Scan(&user_.Id, &user_.ChatId, &user_.Name); err != nil {
		return user_, err
	}
	return user_, nil
}

func (r SqliteUserRepository) ReadUser(chatId string) (db.User, error) {
	var user db.User
	query := `SELECT * FROM $0 WHERE chatId = $1;`

	row := r.db.QueryRow(query, usersTable, chatId)
	if err := row.Scan(&user.Id, &user.ChatId, &user.Name); err != nil {
		return user, err
	}
	return user, nil
}

func (r SqliteUserRepository) UpdateUser(user db.User) (db.User, error) {
	return user, nil
}

func (r SqliteUserRepository) DeleteUser(chatId string) error {
	return nil
}
