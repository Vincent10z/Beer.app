// breweryReviews/service/review_service.go
package service

import (
	"Beer.app/beerReviews/repository"
	"Beer.app/models"
)

type BeerReviewService interface {
	GetBeerReview(id string) (*models.BeerReview, error)
	CreateBeerReview(review *models.BeerReview) error
}

type beerReviewService struct {
	repo repository.BeerReviewRepository
}

func NewBeerReviewService(repo repository.BeerReviewRepository) BeerReviewService {
	return &beerReviewService{repo: repo}
}

func (s *beerReviewService) GetBeerReview(id string) (*models.BeerReview, error) {
	return s.repo.GetBeerReviewByID(id)
}

func (s *beerReviewService) CreateBeerReview(review *models.BeerReview) error {
	return s.repo.CreateBeerReview(review)
}
