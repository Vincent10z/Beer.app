// users/repository/breweryReviews_repository.go
package repository

import (
	"Beer.app/models"
	"errors"
)

type BreweryReviewRepository interface {
	GetBreweryReviewByID(id int) (*models.Review, error)
	CreateBreweryReview(review *models.Review) error
}

type breweryReviewRepository struct {
	breweryReviews map[int]*models.Review
}

func NewBreweryReviewRepository() BreweryReviewRepository {
	return &breweryReviewRepository{breweryReviews: make(map[int]*models.Review)}
}

func (r *breweryReviewRepository) GetBreweryReviewByID(id int) (*models.Review, error) {
	review, exists := r.breweryReviews[id]
	if !exists {
		return nil, errors.New("review not found")
	}
	return review, nil
}

func (r *breweryReviewRepository) CreateBreweryReview(review *models.Review) error {
	if _, exists := r.breweryReviews[review.ID]; exists {
		return errors.New("review with the same ID already exists")
	}
	r.breweryReviews[review.ID] = review
	return nil
}
