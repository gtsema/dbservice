package repository

import (
	"database/sql"
	"dbservice/internal/models/db"
)

type UserRepository interface {
	CreateUser(user db.User) (int, error)
	GetUser(chatId string) (db.User, error)
	DeleteUser(chatId string) error
}

type Repository struct {
	UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
	}
}
