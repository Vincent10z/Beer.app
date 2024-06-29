// users/service/reviews_service.go
package service

import (
	"Beer.app/models"
	"Beer.app/users/repository"
)

type UserService interface {
	GetUser(id int) (*models.User, error)
	CreateUser(user *models.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}
