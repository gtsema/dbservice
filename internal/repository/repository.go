package repository

import (
	"database/sql"
	"dbservice/internal/models/db"
)

type UserRepository interface {
	CreateUser(user db.User) (db.User, error)
	ReadUser(chatId string) (db.User, error)
	UpdateUser(user db.User) (db.User, error)
	DeleteUser(chatId string) error
}

type Repository struct {
	UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{UserRepository: newSqliteUserRepository(db)}
}
