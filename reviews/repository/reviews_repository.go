// users/repository/reviews_repository.go
package repository

import (
	"Beer.app/models"
	"errors"
)

type ReviewRepository interface {
	GetReviewByID(id int) (*models.Review, error)
	CreateReview(review *models.Review) error
}

type reviewRepository struct {
	reviews map[int]*models.Review
}

func NewReviewRepository() ReviewRepository {
	return &reviewRepository{reviews: make(map[int]*models.Review)}
}

func (r *reviewRepository) GetReviewByID(id int) (*models.Review, error) {
	review, exists := r.reviews[id]
	if !exists {
		return nil, errors.New("review not found")
	}
	return review, nil
}

func (r *reviewRepository) CreateReview(review *models.Review) error {
	r.reviews[review.ID] = review
	return nil
}
