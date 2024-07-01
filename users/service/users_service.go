// users/service/breweryReviews_service.go
package service

import (
	"github.com/rotisserie/eris"
	"time"

	"Beer.app/models"
	"Beer.app/users/repository"
	"Beer.app/utils"
)

type UserService interface {
	GetUser(id string) (*models.User, error)
	CreateUser(user *models.User) error
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{UserRepository: repo}
}

// CreateUser Creates a new user
func (s *userService) CreateUser(user *models.User) error {
	user.ID = utils.GenerateUserID()

	hashedPassword, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		return eris.Wrapf(err, "failed to hash password")
	}

	user.PasswordHash = hashedPassword

	user.Account = &models.Account{
		ID:            utils.GenerateAccountID(),
		UserID:        user.ID,
		PlanID:        models.BasicPlan,
		StartDate:     time.Now(),
		EndDate:       time.Now().AddDate(1, 0, 0),
		Status:        "active",
		PaymentMethod: "credit_card",
	}

	if err := s.UserRepository.CreateUser(user); err != nil {
		return eris.Wrapf(err, "failed to create user")
	}

	return nil
}

// GetUser Retrieve a user by ID
func (s *userService) GetUser(id string) (*models.User, error) {
	return s.UserRepository.GetUserByID(id)
}
