// users/repository/brewery_repository.go
package repository

import (
	"Beer.app/models"
	"errors"
)

type BreweryRepository interface {
	GetBreweryByID(id int) (*models.Brewery, error)
	CreateBrewery(brewery *models.Brewery) error
}

type breweryRepository struct {
	breweries map[int]*models.Brewery
}

func NewBreweryRepository() BreweryRepository {
	return &breweryRepository{breweries: make(map[int]*models.Brewery)}
}

func (r *breweryRepository) GetBreweryByID(id int) (*models.Brewery, error) {
	brewery, exists := r.breweries[id]
	if !exists {
		return nil, errors.New("brewery not found")
	}
	return brewery, nil
}

func (r *breweryRepository) CreateBrewery(brewery *models.Brewery) error {
	r.breweries[brewery.ID] = brewery
	return nil
}
