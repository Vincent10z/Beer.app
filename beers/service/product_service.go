// beers/service/beer_service.go
package service

import (
	"Beer.app/beers/repository"
	"Beer.app/models"
)

type BeerService interface {
	GetBeer(id int) (*models.Beer, error)
	CreateBeer(beer *models.Beer) error
}

type beerService struct {
	repo repository.BeerRepository
}

func NewBeerService(repo repository.BeerRepository) BeerService {
	return &beerService{repo: repo}
}

func (s *beerService) GetBeer(id int) (*models.Beer, error) {
	return s.repo.GetBeerByID(id)
}

func (s *beerService) CreateBeer(beer *models.Beer) error {
	return s.repo.CreateBeer(beer)
}
