// users/repository/breweryReviews_repository.go
package repository

import (
	"Beer.app/models"
	"errors"
)

type BeerReviewRepository interface {
	GetBeerReviewByID(id int) (*models.Review, error)
	CreateBeerReview(review *models.Review) error
}

type beerReviewRepository struct {
	reviews map[int]*models.Review
}

func NewBeerReviewRepository() BeerReviewRepository {
	return &beerReviewRepository{reviews: make(map[int]*models.Review)}
}

func (r *beerReviewRepository) GetBeerReviewByID(id int) (*models.Review, error) {
	review, exists := r.reviews[id]
	if !exists {
		return nil, errors.New("review not found")
	}
	return review, nil
}

func (r *beerReviewRepository) CreateBeerReview(review *models.Review) error {
	r.reviews[review.ID] = review
	return nil
}
