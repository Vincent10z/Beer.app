// beerStyles/service/beer_style_service.go
package service

import (
	"Beer.app/beerStyles/repository"
	"Beer.app/models"
	"Beer.app/utils"
	"fmt"
	"github.com/rotisserie/eris"
)

type BeerStyleService interface {
	GetBeerStyle(id string) (*models.BeerStyle, error)
	CreateBeerStyle(style *models.BeerStyle) error
	UpdateBeerStyle(style *models.BeerStyle) error
	DeleteBeerStyle(id string) error
	ListBeerStyles() ([]*models.BeerStyle, error)
	CreateMultipleBeerStyles(styles []*models.BeerStyle) error
}

type beerStyleService struct {
	repo repository.BeerStyleRepository
}

func NewBeerStyleService(repo repository.BeerStyleRepository) BeerStyleService {
	return &beerStyleService{repo: repo}
}

func (s *beerStyleService) GetBeerStyle(id string) (*models.BeerStyle, error) {
	return s.repo.GetBeerStyleByID(id)
}

func (s *beerStyleService) CreateBeerStyle(style *models.BeerStyle) error {
	fmt.Printf("Before ID generation: %s\n", style.ID)
	style.ID = utils.GenerateBeerStyleID()
	fmt.Printf("After ID generation: %s\n", style.ID)

	if err := s.repo.CreateBeerStyle(style); err != nil {
		return eris.Wrapf(err, "failed to create beer style")
	}

	fmt.Printf("After repository creation: %s\n", style.ID)
	return nil
}

func (s *beerStyleService) UpdateBeerStyle(style *models.BeerStyle) error {
	return s.repo.UpdateBeerStyle(style)
}

func (s *beerStyleService) DeleteBeerStyle(id string) error {
	return s.repo.DeleteBeerStyle(id)
}

func (s *beerStyleService) ListBeerStyles() ([]*models.BeerStyle, error) {
	return s.repo.ListBeerStyles()
}

func (s *beerStyleService) CreateMultipleBeerStyles(styles []*models.BeerStyle) error {
	for _, style := range styles {
		style.ID = utils.GenerateBeerStyleID()
	}

	if err := s.repo.CreateMultipleBeerStyles(styles); err != nil {
		return eris.Wrapf(err, "failed to create multiple beer styles")
	}

	return nil
}
