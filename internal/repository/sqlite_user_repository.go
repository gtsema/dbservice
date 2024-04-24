package repository

import (
	"database/sql"
	"dbservice/internal/models/db"
	"errors"
)

type SqliteUserRepository struct {
	db *sql.DB
}

func newSqliteUserRepository(database *sql.DB) *SqliteUserRepository {
	return &SqliteUserRepository{db: database}
}

func (r SqliteUserRepository) CreateUser(user db.User) (db.User, error) {
	var user_ db.User
	query := "INSERT INTO users (chatId, name, hash, salt) VALUES ($1, $2, $3, $4) RETURNING id, chatId, name, hash, salt;"

	row := r.db.QueryRow(query, user.ChatId, user.Name, user.Hash, user.Salt)
	if err := row.Scan(&user_.Id, &user_.ChatId, &user_.Name, &user_.Hash, &user_.Salt); err != nil {
		return user_, err
	}
	return user_, nil
}

func (r SqliteUserRepository) ReadUser(chatId string) (db.User, error) {
	var user_ db.User
	query := "SELECT id, chatId, name, hash, salt FROM users WHERE chatId = $1;"

	row := r.db.QueryRow(query, chatId)
	if err := row.Scan(&user_.Id, &user_.ChatId, &user_.Name, &user_.Hash, &user_.Salt); err != nil {
		return user_, err
	}
	return user_, nil
}

func (r SqliteUserRepository) UpdateUser(user db.User) (db.User, error) {
	var user_ db.User
	query := "UPDATE users SET chatId = $1, name = $2, hash = $3, salt = $4 WHERE id = $5 RETURNING id, chatId, name, hash, salt;"

	row := r.db.QueryRow(query, user.ChatId, user.Name, user.Hash, user.Salt, user.Id)
	if err := row.Scan(&user_.Id, &user_.ChatId, &user_.Name, &user_.Hash, &user_.Salt); err != nil {
		return user_, err
	}
	return user_, nil
}

func (r SqliteUserRepository) DeleteUser(chatId string) error {
	query := "DELETE FROM users WHERE chatId = $1;"

	res, err := r.db.Exec(query, chatId)
	if err != nil {
		return err
	}

	numDeleted, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if numDeleted == 0 {
		return errors.New("user not found")
	}

	return nil
}
