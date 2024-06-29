// users/service/brewery_service.go
package service

import (
	"Beer.app/models"
	"Beer.app/users/repository"
)

type BreweryService interface {
	GetUser(id int) (*models.Brewery, error)
	CreateUser(user *models.Brewery) error
}

type breweryService struct {
	repo repository.BreweryRepository
}

func NewBreweryService(repo repository.BreweryRepository) BreweryService {
	return &breweryService{repo: repo}
}

func (s *breweryService) GetBrewery(id int) (*models.Brewery, error) {
	return s.repo.GetBreweryByID(id)
}

func (s *breweryService) CreateBrewery(user *models.Brewery) error {
	return s.repo.CreateBrewery(user)
}
