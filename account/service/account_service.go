// accounts/service/account_service.go
package service

import (
	"Beer.app/account/repository"
	"Beer.app/models"
)

type AccountService interface {
	GetAccount(id string) (*models.Account, error)
	CreateAccount(account *models.Account) error
}

type accountService struct {
	AccountRepository repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{AccountRepository: repo}
}

func (s *accountService) GetAccount(id string) (*models.Account, error) {
	return s.AccountRepository.GetAccountByID(id)
}

func (s *accountService) CreateAccount(account *models.Account) error {
	return s.AccountRepository.CreateAccount(account)
}
