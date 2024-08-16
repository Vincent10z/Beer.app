// beers/repository/beer_repository.go
package repository

import (
	"Beer.app/models"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type beerRepository struct {
	db *gorm.DB
}

type BeerRepository interface {
	GetBeerByID(id string) (*models.Beer, error)
	CreateBeer(beer *models.Beer) error
}

func NewBeerRepository(db *gorm.DB) BeerRepository {
	return &beerRepository{db: db}
}

func (r *beerRepository) GetBeerByID(id string) (*models.Beer, error) {
	beer := &models.Beer{}

	if err := r.db.Where("id = ?", id).First(beer).Error; err != nil {
		return nil, eris.Wrapf(err, "failed to get beer with id %d", id)
	}

	return beer, nil
}

func (r *beerRepository) CreateBeer(beer *models.Beer) error {

	if err := r.db.Create(beer).Error; err != nil {
		return eris.Wrapf(err, "failed to create beer")
	}
	return nil
}
