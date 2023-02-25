package service

import "github.com/DanielHernandezO/bookstore_users-api/domain"

type UserService interface {
	CreateUser(user domain.User) (*domain.User, error)
}

type userService struct {
}

func NewUserService() *userService {
	return &userService{}
}
func (u userService) CreateUser(user domain.User) (*domain.User, error) {
	return nil, nil
}
