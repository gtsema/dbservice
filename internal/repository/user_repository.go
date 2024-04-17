package repository

import (
	"database/sql"
	"dbservice/internal/models/db"
	"fmt"
)

type UserRepositorySqlite struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositorySqlite {
	return &UserRepositorySqlite{db: db}
}

func (r *UserRepositorySqlite) CreateUser(user db.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (chatId, name, hash, salt) VALUES ('$1', '$2', '$3', '$4')", usersTable)

	row := r.db.QueryRow(query, user.ChatId, user.Name, user.Hash, user.Salt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepositorySqlite) GetUser(chatId string) (db.User, error) {
	var user db.User
	query := fmt.Sprintf("SELECT id, chatId, name FROM %s WHERE chatId = '$1'", usersTable)

	row := r.db.QueryRow(query, chatId)
	err := row.Scan(&user.Id, &user.ChatId, &user.Name)

	return user, err
}

func (r *UserRepositorySqlite) DeleteUser(chatId string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE chatId = '$1'", usersTable)

	_, err := r.db.Exec(query, chatId)
	return err
}
