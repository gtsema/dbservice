package service

import (
	"dbservice/internal/models/api"
	"dbservice/internal/models/db"
	"dbservice/internal/repository"
)

type UserServiceRest struct {
	repository repository.UserRepository
}

func NewUserServiceRest(repository repository.UserRepository) *UserServiceRest {
	return &UserServiceRest{repository: repository}
}

func (s *UserServiceRest) CreateUser(user *api.User) (int, error) {
	id, err := s.repository.CreateUser(db.User{
		ChatId: user.ChatId,
		Name:   user.Name,
		Hash:   "0",
		Salt:   "cats",
	})

	return id, err
}

func (s *UserServiceRest) GetUser(chatId string) (*api.User, error) {
	user, err := s.repository.GetUser(chatId)
	if err != nil {
		return nil, err
	}

	return &api.User{
		Id:     user.Id,
		ChatId: user.ChatId,
		Name:   user.Name,
	}, nil
}

func (s *UserServiceRest) DeleteUser(chatId string) error {
	err := s.repository.DeleteUser(chatId)
	return err
}
