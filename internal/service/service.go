package service

import (
	"dbservice/internal/models/api"
	"dbservice/internal/models/db"
	"dbservice/internal/repository"
)

type UserService interface {
	CreateUser(user *api.User) (*db.User, error)
	GetUser(chatId string) (*api.User, error)
	DeleteUser(ChatId string) error
}

type Service struct {
	UserService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		UserService: NewUserServiceRest(r),
	}
}
