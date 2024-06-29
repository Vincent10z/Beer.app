// users/repository/breweryReviews_repository.go
package repository

import (
	"Beer.app/models"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type BeerReviewRepository interface {
	GetBeerReviewByID(id int) (*models.Review, error)
	CreateBeerReview(review *models.Review) error
}

type beerReviewRepository struct {
	db *gorm.DB
}

func NewBeerReviewRepository(db *gorm.DB) BeerReviewRepository {
	return &beerReviewRepository{db: db}
}

func (r *beerReviewRepository) GetBeerReviewByID(id int) (*models.Review, error) {
	beerReview := &models.Review{}

	if err := r.db.Where("id = ?", id).First(beerReview).Error; err != nil {
		return nil, eris.Wrapf(err, "failed to get beer review with id %d", id)
	}

	return beerReview, nil
}

func (r *beerReviewRepository) CreateBeerReview(review *models.Review) error {
	if err := r.db.Create(review).Error; err != nil {
		return eris.Wrap(err, "failed to create beer review")
	}

	return nil
}
