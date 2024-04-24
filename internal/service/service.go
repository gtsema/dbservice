package service

import (
	"dbservice/internal/models/api"
	"dbservice/internal/repository"
)

type UserService interface {
	CreateUser(user *api.User) (*api.User, error)
	GetUser(chatId string) (*api.User, error)
	UpdateUser(user *api.User) (*api.User, error)
	DeleteUser(chatId string) error
}

type Service struct {
	UserService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		UserService: NewUserServiceRest(r),
	}
}
