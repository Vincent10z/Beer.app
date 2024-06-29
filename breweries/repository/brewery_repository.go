// users/repository/breweryReviews_repository.go
package repository

import (
	"Beer.app/models"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type breweryRepository struct {
	db *gorm.DB
}

type BreweryRepository interface {
	GetBreweryByID(id int) (*models.Brewery, error)
	CreateBrewery(brewery *models.Brewery) error
}

func NewBreweryRepository(db *gorm.DB) BreweryRepository {
	return &breweryRepository{db: db}
}

func (r *breweryRepository) GetBreweryByID(id int) (*models.Brewery, error) {
	brewery := &models.Brewery{}

	if err := r.db.Where("id = ?", id).First(brewery).Error; err != nil {
		return nil, eris.Wrapf(err, "failed to get brewery with id %d", id)
	}

	return brewery, nil
}

func (r *breweryRepository) CreateBrewery(brewery *models.Brewery) error {
	if err := r.db.Create(brewery).Error; err != nil {
		return eris.Wrap(err, "failed to create brewery")
	}

	return nil
}
