// breweryReviews/service/review_service.go
package service

import (
	"Beer.app/breweryReviews/repository"
	"Beer.app/models"
)

type BreweryReviewService interface {
	GetBreweryReview(id string) (*models.BreweryReview, error)
	CreateBreweryReview(review *models.BreweryReview) error
}

type breweryReviewService struct {
	repo repository.BreweryReviewRepository
}

func NewReviewService(repo repository.BreweryReviewRepository) BreweryReviewService {
	return &breweryReviewService{repo: repo}
}

func (s *breweryReviewService) GetBreweryReview(id string) (*models.BreweryReview, error) {
	return s.repo.GetBreweryReviewByID(id)
}

func (s *breweryReviewService) CreateBreweryReview(review *models.BreweryReview) error {
	return s.repo.CreateBreweryReview(review)
}
