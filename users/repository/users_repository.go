// users/repository/breweryReviews_repository.go
package repository

import (
	"Beer.app/models"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}

	if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, eris.Wrapf(err, "failed to get user with id %d", id)
	}

	return user, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return eris.Wrap(err, "failed to create user")
	}

	return nil
}
