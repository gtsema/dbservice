package repository

import (
	"dbservice/pkg/domain"
)

type UserRepository interface {
	FindByChatId(chatId int) (*domain.User, error)
	Create(user *domain.User) error
	Delete(user *domain.User) error
}
