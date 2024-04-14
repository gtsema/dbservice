package sqlite

import (
	"database/sql"
	"dbservice/pkg/domain"
)

type SqliteUserRepository struct {
	db *sql.DB
}

func (r *SqliteUserRepository) FindByChatId(chatId int) (*domain.User, error) {
	var user domain.User
	return &user, r.db.
		QueryRow("SELECT 'Id', 'ChatId', 'Name', 'Hash', 'Salt' FROM User WHERE 'ChatId' = ?", chatId).
		Scan(&user.Id, &user.ChatId, &user.Name, &user.Hash, &user.Salt)
}

func (r *SqliteUserRepository) Create(user *domain.User) error {
	_, err := r.db.Exec(
		"INSERT INTO USER (ChatId, Name, Hash, Salt) VALUES (?, ?, ?, ?)",
		user.ChatId, user.Name, user.Hash, user.Salt)
	return err
}

func (r *SqliteUserRepository) Delete(user *domain.User) error {
	_, err := r.db.Exec("DELETE FROM USER WHERE ChatId = ?", user.ChatId)
	return err
}
