// beers/repository/beer_repository.go
package repository

import (
	"Beer.app/models"
	"errors"
)

type BeerRepository interface {
	GetBeerByID(id int) (*models.Beer, error)
	CreateBeer(beer *models.Beer) error
}

type beerRepository struct {
	beers map[int]*models.Beer
}

func NewBeerRepository() BeerRepository {
	return &beerRepository{beers: make(map[int]*models.Beer)}
}

func (r *beerRepository) GetBeerByID(id int) (*models.Beer, error) {
	beer, exists := r.beers[id]
	if !exists {
		return nil, errors.New("beer not found")
	}
	return beer, nil
}

func (r *beerRepository) CreateBeer(beer *models.Beer) error {
	r.beers[beer.Id] = beer
	return nil
}
