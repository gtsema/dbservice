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

func (s *UserServiceRest) CreateUser(user *api.User) (*api.User, error) {
	user_, err := s.repository.CreateUser(db.User{
		ChatId: user.ChatId,
		Name:   user.Name,
		Hash:   "cats",
		Salt:   "cats",
	})

	return &api.User{
		Id:     user_.Id,
		Name:   user_.Name,
		ChatId: user_.ChatId,
	}, err
}

func (s *UserServiceRest) GetUser(chatId string) (*api.User, error) {
	user, err := s.repository.ReadUser(chatId)
	if err != nil {
		return nil, err
	}

	return &api.User{
		Id:     user.Id,
		ChatId: user.ChatId,
		Name:   user.Name,
	}, nil
}

func (s *UserServiceRest) UpdateUser(user *api.User) (*api.User, error) {
	user_, err := s.repository.UpdateUser(db.User{
		Id:     user.Id,
		ChatId: user.ChatId,
		Name:   user.Name,
		Hash:   "cats",
		Salt:   "cats",
	})

	return &api.User{
		Id:     user_.Id,
		Name:   user_.Name,
		ChatId: user_.ChatId,
	}, err
}

func (s *UserServiceRest) DeleteUser(chatId string) error {
	err := s.repository.DeleteUser(chatId)
	return err
}
