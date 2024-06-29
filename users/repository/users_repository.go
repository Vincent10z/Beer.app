// users/repository/breweryReviews_repository.go
package repository

import (
	"Beer.app/models"
	"errors"
)

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) error
}

type userRepository struct {
	users map[int]*models.User
}

func NewUserRepository() UserRepository {
	return &userRepository{users: make(map[int]*models.User)}
}

func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
	r.users[user.AccountId] = user
	return nil
}
