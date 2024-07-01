// users/repository/account_repository.go
package repository

import (
	"Beer.app/models"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAccountByID(id string) (*models.Account, error)
	CreateAccount(account *models.Account) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) GetAccountByID(id string) (*models.Account, error) {
	account := &models.Account{}

	if err := r.db.Where("id = ?", id).First(account).Error; err != nil {
		return nil, eris.Wrapf(err, "failed to get account with id %d", id)
	}

	return account, nil
}

func (r *accountRepository) CreateAccount(account *models.Account) error {
	if err := r.db.Create(account).Error; err != nil {
		return eris.Wrap(err, "failed to create account")
	}

	return nil
}
