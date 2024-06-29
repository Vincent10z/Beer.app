// users/repository/breweryReviews_repository.go
package repository

import (
	"Beer.app/models"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type BreweryReviewRepository interface {
	GetBreweryReviewByID(id int) (*models.Review, error)
	CreateBreweryReview(review *models.Review) error
}

type breweryReviewRepository struct {
	db *gorm.DB
}

func NewBreweryReviewRepository(db *gorm.DB) BreweryReviewRepository {
	return &breweryReviewRepository{db: db}
}

func (r *breweryReviewRepository) GetBreweryReviewByID(id int) (*models.Review, error) {
	review := &models.Review{}

	if err := r.db.Where("id = ?", id).First(review).Error; err != nil {
		return nil, eris.Wrapf(err, "failed to get brewery review with id %d", id)
	}

	return review, nil
}

func (r *breweryReviewRepository) CreateBreweryReview(review *models.Review) error {
	if err := r.db.Create(review).Error; err != nil {
		return eris.Wrap(err, "failed to create brewery review")
	}

	return nil
}
