package service

import (
	"dbservice/pkg/domain"
	"dbservice/pkg/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func (s *UserService) GetUser(chatId int) (*domain.User, error) {
	return s.repository.FindByChatId(chatId)
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.repository.Create(user)
}

func (s *UserService) DeleteUser(user *domain.User) error {
	return s.repository.Delete(user)
}
