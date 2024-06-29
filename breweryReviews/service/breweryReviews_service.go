// breweryReviews/service/review_service.go
package service

import (
	"Beer.app/breweryReviews/repository"
	"Beer.app/models"
)

type ReviewService interface {
	GetReview(id int) (*models.Review, error)
	CreateReview(review *models.Review) error
}

type reviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) ReviewService {
	return &reviewService{repo: repo}
}

func (s *reviewService) GetReview(id int) (*models.Review, error) {
	return s.repo.GetReviewByID(id)
}

func (s *reviewService) CreateReview(review *models.Review) error {
	return s.repo.CreateReview(review)
}
