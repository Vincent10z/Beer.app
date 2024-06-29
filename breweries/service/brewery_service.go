// users/service/breweryReviews_service.go
package service

import (
	"Beer.app/breweries/repository"
	"Beer.app/models"
)

type BreweryService interface {
	GetBrewery(id int) (*models.Brewery, error)
	CreateBrewery(user *models.Brewery) error
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
