// beerStyles/repository/beer_style_repository.go
package repository

import (
	"Beer.app/models"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type BeerStyleRepository interface {
	GetBeerStyleByID(id string) (*models.BeerStyle, error)
	CreateBeerStyle(style *models.BeerStyle) error
	UpdateBeerStyle(style *models.BeerStyle) error
	DeleteBeerStyle(id string) error
	ListBeerStyles() ([]*models.BeerStyle, error)
	CreateMultipleBeerStyles(styles []*models.BeerStyle) error
}

type beerStyleRepository struct {
	db *gorm.DB
}

func NewBeerStyleRepository(db *gorm.DB) BeerStyleRepository {
	return &beerStyleRepository{db: db}
}

func (r *beerStyleRepository) GetBeerStyleByID(id string) (*models.BeerStyle, error) {
	var style models.BeerStyle
	if err := r.db.First(&style, "id = ?", id).Error; err != nil {
		return nil, eris.Wrapf(err, "failed to get beer style with id %s", id)
	}
	return &style, nil
}

func (r *beerStyleRepository) CreateBeerStyle(style *models.BeerStyle) error {
	if err := r.db.Create(style).Error; err != nil {
		return eris.Wrap(err, "failed to create beer style")
	}
	return nil
}

func (r *beerStyleRepository) UpdateBeerStyle(style *models.BeerStyle) error {
	if err := r.db.Save(style).Error; err != nil {
		return eris.Wrap(err, "failed to update beer style")
	}
	return nil
}

func (r *beerStyleRepository) DeleteBeerStyle(id string) error {
	if err := r.db.Delete(&models.BeerStyle{}, "id = ?", id).Error; err != nil {
		return eris.Wrapf(err, "failed to delete beer style with id %s", id)
	}
	return nil
}

func (r *beerStyleRepository) ListBeerStyles() ([]*models.BeerStyle, error) {
	var styles []*models.BeerStyle
	if err := r.db.Find(&styles).Error; err != nil {
		return nil, eris.Wrap(err, "failed to list beer styles")
	}
	return styles, nil
}

func (r *beerStyleRepository) CreateMultipleBeerStyles(styles []*models.BeerStyle) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, style := range styles {
			if err := tx.Create(style).Error; err != nil {
				return eris.Wrap(err, "failed to create beer style")
			}
		}
		return nil
	})
	if err != nil {
		return eris.Wrap(err, "failed to create multiple beer styles")
	}
	return nil
}
