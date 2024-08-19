package repository

import (
	"Beer.app/models"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type BeerRepository interface {
	CreateBeer(beer *models.Beer) (*models.Beer, error)
	GetBeerByID(id string) (*models.Beer, error)
	UpdateBeer(beer *models.Beer) (*models.Beer, error)
	DeleteBeer(id string) error
	ListBeers() ([]*models.Beer, error)
}

type beerRepository struct {
	db *gorm.DB
}

func NewBeerRepository(db *gorm.DB) BeerRepository {
	return &beerRepository{db: db}
}

func (r *beerRepository) CreateBeer(beer *models.Beer) (*models.Beer, error) {
	if err := r.db.Create(beer).Error; err != nil {
		return nil, eris.Wrap(err, "failed to create beer")
	}
	return beer, nil
}

func (r *beerRepository) GetBeerByID(id string) (*models.Beer, error) {
	var beer models.Beer
	if err := r.db.First(&beer, "id = ?", id).Error; err != nil {
		return nil, eris.Wrapf(err, "failed to get beer with id %s", id)
	}
	return &beer, nil
}

func (r *beerRepository) UpdateBeer(beer *models.Beer) (*models.Beer, error) {
	if err := r.db.Save(beer).Error; err != nil {
		return nil, eris.Wrap(err, "failed to update beer")
	}
	return beer, nil
}

func (r *beerRepository) DeleteBeer(id string) error {
	if err := r.db.Delete(&models.Beer{}, "id = ?", id).Error; err != nil {
		return eris.Wrapf(err, "failed to delete beer with id %s", id)
	}
	return nil
}

func (r *beerRepository) ListBeers() ([]*models.Beer, error) {
	var beers []*models.Beer
	if err := r.db.Find(&beers).Error; err != nil {
		return nil, eris.Wrap(err, "failed to list beers")
	}
	return beers, nil
}
