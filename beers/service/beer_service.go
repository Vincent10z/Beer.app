package service

import (
	"Beer.app/beers/repository"
	"Beer.app/models"
)

type BeerService interface {
	GetBeer(id string) (*models.Beer, error)
	CreateBeer(beer *models.Beer) (*models.Beer, error)
	UpdateBeer(beer *models.Beer) (*models.Beer, error)
	DeleteBeer(id string) error
	ListBeers() ([]*models.Beer, error)
}

type beerService struct {
	repo repository.BeerRepository
}

func NewBeerService(repo repository.BeerRepository) BeerService {
	return &beerService{repo: repo}
}

func (s *beerService) GetBeer(id string) (*models.Beer, error) {
	return s.repo.GetBeerByID(id)
}

func (s *beerService) CreateBeer(beer *models.Beer) (*models.Beer, error) {
	return s.repo.CreateBeer(beer)
}

func (s *beerService) UpdateBeer(beer *models.Beer) (*models.Beer, error) {
	return s.repo.UpdateBeer(beer)
}

func (s *beerService) DeleteBeer(id string) error {
	return s.repo.DeleteBeer(id)
}

func (s *beerService) ListBeers() ([]*models.Beer, error) {
	return s.repo.ListBeers()
}
